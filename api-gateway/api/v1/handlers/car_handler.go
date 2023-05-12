package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/dailoi280702/se121/api-gateway/internal/utils"
	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	// :TODO change this?? in middle function??
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// get request data from api call
		req := &car.GetCarReq{}
		err := utils.DecodeJSONBody(w, r, req)
		if err != nil {
			var mr *utils.MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				MustSendError(err, w)
			}
			return
		}

		// send request to service
		res, err := h.carService.GetCar(context.Background(), req)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			code := status.Code(err)
			s, _ := status.FromError(err)

			switch code {
			case codes.InvalidArgument:
				w.WriteHeader(http.StatusBadRequest)
			case codes.AlreadyExists:
				w.WriteHeader(http.StatusConflict)
			case codes.Unavailable:
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			default:
				MustSendError(err, w)
				return
			}

			_, err = w.Write([]byte(s.Message()))
			if err != nil {
				MustSendError(err, w)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Panic(err)
		}
	})

	return r
}
