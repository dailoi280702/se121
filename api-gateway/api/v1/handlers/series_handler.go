package handlers

import (
	"context"
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
		var req car.GetSeriesReq
		var res *car.Series
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.GetSeries(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleUpdateSeries(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.UpdateSeriesReq
		convertJsonApiToGrpc(w, r,
			func() error {
				_, err := carService.UpdateSeries(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleCreateSeries(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.CreateSeriesReq
		convertJsonApiToGrpc(w, r,
			func() error {
				_, err := carService.CreateSeries(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleSearchSeries(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.SearchReq
		var res *car.SearchForSeriesRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.SearchForSeries(context.Background(), &req)
				return err
			},
			convertWithUrlQuery(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}
