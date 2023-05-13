package handlers

import (
	"context"
	"net/http"

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

// get /
// put /
// post /
// delete /
// get / brand /
// get /

func (h *CarHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	})

	return r
}

func handleGetCar(w http.ResponseWriter, r *http.Request, carService car.CarServiceClient) {
	var req car.GetCarReq
	var car *car.Car
	convertJsonApiToGrpc(w, r, func() error {
		var err error
		car, err = carService.GetCar(context.Background(), &req)
		return err
	},
		convertWithJsonReqData(&req),
		convertWithPostFunc(func() {
			SendJson(w, car)
		}))
}

func handleCreateCar(w http.ResponseWriter, r *http.Request, carService car.CarServiceClient) {
	var req car.CreateCarReq
	convertJsonApiToGrpc(w, r, func() error {
		var err error
		_, err = carService.CreateCar(context.Background(), &req)
		return err
	}, convertWithJsonReqData(&req))
}

func handleUpdateCar(w http.ResponseWriter, r *http.Request, carService car.CarServiceClient) {
	var req car.UpdateCarReq
	convertJsonApiToGrpc(w, r, func() error {
		var err error
		_, err = carService.UpdateCar(context.Background(), &req)
		return err
	}, convertWithJsonReqData(&req))
}

func handleSearchCar(w http.ResponseWriter, r *http.Request, carService car.CarServiceClient) {
	var req car.SearchForCarReq
	convertJsonApiToGrpc(w, r, func() error {
		var err error
		_, err = carService.SearchForCar(context.Background(), &req)
		return err
	}, convertWithJsonReqData(&req))
}
