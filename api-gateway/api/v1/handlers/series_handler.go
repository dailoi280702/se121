package handlers

import (
	"net/http"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/go-chi/chi/v5"
)

func NewSeriesRoutes(carService car.CarServiceClient) chi.Router {
	c := chi.NewRouter()

	c.Get("/", handleGetSeries(carService))
	c.Put("/", handleUpdateSeries(carService))
	c.Post("/", handleCreateSeries(carService))
	c.Get("/search", handleSearchSeries(carService))

	return c
}

func handleGetSeries(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func handleUpdateSeries(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func handleCreateSeries(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func handleSearchSeries(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
