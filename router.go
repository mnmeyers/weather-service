package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"net/http"
)

// GetRouter returns the mux that handles all incoming HTTP Requests
func GetRouter() http.Handler {
	router, c := makeRouter()

	return c.Handler(router)
}

func makeRouter() (*chi.Mux, *cors.Cors) {
	router := chi.NewRouter()
	controller := GetController()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:           300,
		Debug:            false,
	})

	router.Use(c.Handler)
	router.Use(middleware.Recoverer)

	router.Route("/", func(innerRouter chi.Router) {
		innerRouter.Get("/weather/{lat}/{lon}", controller.GetWeather)
	})

	return router, c
}
