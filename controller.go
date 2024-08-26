package main

import (
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

var _ Controller = (*ControllerImpl)(nil)

func GetController() Controller {
	return &ControllerImpl{
		service: GetService(),
	}
}

func (controller *ControllerImpl) GetWeather(w http.ResponseWriter, r *http.Request) {
	latStr := r.URL.Query().Get("lat")
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lonStr := r.URL.Query().Get("lon")
	lon, err := strconv.ParseFloat(lonStr, 64)
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
