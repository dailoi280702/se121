package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
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

func NewCarRoutes(carService car.CarServiceClient) chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", handleGetCarById(carService))
	r.Get("/{id}/related", handleGetRelatedCars(carService))
	r.Get("/", handleGetCars(carService))
	r.Delete("/", handleDeleteCarById(carService))
	r.Put("/", handleUpdateCar(carService))
	r.Post("/", handleCreateCar(carService))

	r.Get("/index", handleGetCarMetaData(carService))
	r.Get("/search", handleSearchCar(carService))

	return r
}

func handleGetRelatedCars(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		req := car.GetRelatedCarReq{Id: int32(id)}
		var res *car.GetRelatedCarRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.GetRelatedCar(context.Background(), &req)
				return err
			},
			convertWithUrlQuery(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleUpdateCar(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.UpdateCarReq
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				_, err = carService.UpdateCar(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleDeleteCarById(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.DeleteCarReq
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				_, err = carService.DeleteCar(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req))
	}
}

func handleGetCarById(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		req := car.GetCarReq{Id: int32(id)}
		var res *car.Car
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.GetCar(context.Background(), &req)
				return err
			},
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleGetCars(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.GetCarsReq
		var res *car.GetCarsRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.GetCars(context.Background(), &req)
				return err
			},
			convertWithUrlQuery(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}

func handleGetCarMetaData(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res *car.GetCarMetadataRes
		convertJsonApiToGrpc(w, r, func() error {
			var err error
			res, err = carService.GetCarMetadata(context.Background(), &utils.Empty{})
			return err
		}, convertWithPostFunc(func() {
			SendJson(w, res)
		}))
	}
}

func handleCreateCar(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req car.CreateCarReq
		var res *car.CreateCarRes
		convertJsonApiToGrpc(
			w, r,
			func() error {
				var err error
				res, err = carService.CreateCar(context.Background(), &req)
				return err
			},
			convertWithJsonReqData(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}),
		)
	}
}

func handleSearchCar(carService car.CarServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req utils.SearchReq
		var res *car.SearchForCarRes
		convertJsonApiToGrpc(w, r,
			func() error {
				var err error
				res, err = carService.SearchForCar(context.Background(), &req)
				return err
			},
			convertWithUrlQuery(&req),
			convertWithPostFunc(func() {
				SendJson(w, res)
			}))
	}
}
