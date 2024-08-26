package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

// Controller defines the interface that must be implemented
type Controller interface {
	GetWeather(w http.ResponseWriter, r *http.Request)
}

type ControllerImpl struct {
	service Service
}

type Gridpoint struct {
	Properties struct {
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

type Forecast struct {
	Properties struct {
		Periods []struct {
			Name          string `json:"name"`
			Temperature   int    `json:"temperature"`
			ShortForecast string `json:"shortForecast"`
		} `json:"periods"`
	} `json:"properties"`
}

var _ Controller = (*ControllerImpl)(nil)

func GetController() Controller {
	return &ControllerImpl{
		service: GetService(),
	}
}

func (controller *ControllerImpl) GetWeather(w http.ResponseWriter, r *http.Request) {
	lat, err := strconv.ParseFloat(chi.URLParam(r, "lat"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lon, err := strconv.ParseFloat(chi.URLParam(r, "lon"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	weather, err := controller.service.GetWeather(r.Context(), lat, lon)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(weather))
}
