package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"github.com/go-chi/chi/v5"
)

func NewBrandRoutes(carService car.CarServiceClient) chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", handleGetBrand(carService))
	r.Post("/", handleCreateBrand(carService))
	r.Put("/", handleUpdateBrand(carService))
	r.Get("/search", handleSearchForBrand(carService))

	return r
}

func handleGetBrand(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		req := car.GetBrandReq{Id: int32(id)}
		var res *car.Brand
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.GetBrand(context.Background(), &req)
				return err
			},
			// convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleCreateBrand(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.CreateBrandReq
		var res *car.CreateBrandRes
		convertJsonApiToGrpc(
			w, r,
			func() error {
				var err error
				res, err = carService.CreateBrand(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}),
		)
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
		var req utils.SearchReq
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
