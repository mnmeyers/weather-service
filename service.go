package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Service interface {
	GetWeather(ctx context.Context, lat, lon float64) (string, error)
}

type ServiceImpl struct{}

var _ Service = (*ServiceImpl)(nil)
var service Service
var onceService sync.Once

// GetService returns a thread-safe singleton of the weather service.
func GetService() Service {
	onceService.Do(func() {
		service = &ServiceImpl{}
	})

	return service
}

// GetWeather returns the weather for the given coordinates
func (s *ServiceImpl) GetWeather(ctx context.Context, lat, lon float64) (string, error) {
	// fetch weather from Weather API
	// Step 1: Get the gridpoint based on latitude and longitude
	gridURL := fmt.Sprintf("https://api.weather.gov/points/%f,%f", lat, lon)
	resp, err := http.Get(gridURL)
	if err != nil {
		fmt.Print("error fetching weather from Weather API: ", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("error reading weather from Weather API: ", err)
		return "", err
	}

	if resp.StatusCode != 200 {
		fmt.Print("error fetching weather from Weather API: ", resp.StatusCode)
		return "", fmt.Errorf("error fetching weather from Weather API: %d", resp.StatusCode)
	}

	var gridpoint Gridpoint
	if err := json.Unmarshal(body, &gridpoint); err != nil {
		fmt.Print("error unmarshalling weather from Weather API: ", err)
		return "", err
	}

	// Step 2: Use the gridpoint to get the weather forecast
	forecastURL := gridpoint.Properties.Forecast
	resp, err = http.Get(forecastURL)
	if err != nil {
		fmt.Print("error fetching forecast from Weather API: ", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("error reading forecast from Weather API: ", err)
		return "", err
	}

	var forecast Forecast
	if err := json.Unmarshal(body, &forecast); err != nil {
		fmt.Print("error unmarshalling forecast from Weather API: ", err)
		return "", err
	}

	// Step 3: Get the short forecast for the current period
	if len(forecast.Properties.Periods) > 0 {
		currForecast := forecast.Properties.Periods[0]
		forecastWithCharacterization := fmt.Sprintf("%s and %s", currForecast.ShortForecast, getWeatherCharacterization(currForecast.Temperature))
		return forecastWithCharacterization, nil
	}

	fmt.Print("no forecast data available")
	return "", fmt.Errorf("no forecast data available")
}

func getWeatherCharacterization(temp int) string {
	if temp > 90 {
		return "hot"
	} else if temp < 40 {
		return "cold"
	}
	return "moderate"
}
