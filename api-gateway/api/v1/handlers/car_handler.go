package handlers

import (
	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/go-chi/chi/v5"
)

type CarHandler struct {
	carService car.CarServiceClient
}

func NewCarHandler(carService car.CarServiceClient) *CarHandler {
	return &CarHandler{
		carService: carService,
	}
}

func (h *CarHandler) Routes() chi.Router {
	r := chi.NewRouter()

	// :TODO define routes

	return r
}
