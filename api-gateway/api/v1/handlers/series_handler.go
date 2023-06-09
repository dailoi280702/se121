package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"github.com/go-chi/chi/v5"
)

func NewSeriesRoutes(carService car.CarServiceClient) chi.Router {
	c := chi.NewRouter()

	c.Get("/{id}", handleGetSeries(carService))
	c.Get("/", handleGetAllSeries(carService))
	c.Put("/", handleUpdateSeries(carService))
	c.Post("/", handleCreateSeries(carService))
	c.Get("/search", handleSearchSeries(carService))

	return c
}

func handleGetSeries(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		req := car.GetSeriesReq{Id: int32(id)}
		var res *car.Series
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.GetSeries(context.Background(), &req)
				return err
			},
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleGetAllSeries(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.GetAllSeriesReq
		var res *car.GetAllSeriesRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.GetAllSeries(context.Background(), &req)
				return err
			},
			convertWithUrlQuery(&req),
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
		var res *car.CreateSeriesRes
		convertJsonApiToGrpc(
			w, r,
			func() error {
				var err error
				res, err = carService.CreateSeries(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}),
		)
	}
}

func handleSearchSeries(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req utils.SearchReq
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
