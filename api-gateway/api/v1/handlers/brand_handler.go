package handlers

import (
	"context"
	"net/http"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/go-chi/chi/v5"
)

func NewBrandRoutes(carService car.CarServiceClient) chi.Router {
	r := chi.NewRouter()

	r.Get("/", handleGetBrand(carService))
	r.Post("/", handleCreateBrand(carService))
	r.Put("/", handleUpdateBrand(carService))
	r.Get("/search", handleSearchForBrand(carService))

	return r
}

func handleGetBrand(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.GetBrandReq
		var res *car.Brand
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.GetBrand(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleCreateBrand(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.CreateBrandReq
		convertJsonApiToGrpc(w, r,
			func() error {
				_, err := carService.CreateBrand(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleUpdateBrand(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.UpdateBrandReq
		convertJsonApiToGrpc(w, r,
			func() error {
				_, err := carService.UpdateBrand(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleSearchForBrand(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.SearchReq
		var res *car.SearchForBrandRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.SearchForBrand(context.Background(), &req)
				return err
			},
			convertWithUrlQuery(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}
