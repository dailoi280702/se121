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
