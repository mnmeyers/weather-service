package main

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
