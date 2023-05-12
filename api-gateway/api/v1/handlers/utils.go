package handlers

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MustSendError(err error, w http.ResponseWriter) {
	log.Printf("\nBitch, you got an err: %s\n", err.Error())
	http.Error(w, fmt.Sprintf("server errror: %s", err.Error()), http.StatusInternalServerError)
}

func GprcToHttp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func SendJsonFromGrpcError(w http.ResponseWriter, err error, actions *map[codes.Code]func()) {
	code := status.Code(err)
	s, _ := status.FromError(err)

	ok := false
	var action func()
	if actions != nil {
		action, ok = (*actions)[code]
	}

	switch {
	case ok:
		action()
	case code == codes.InvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	case code == codes.AlreadyExists:
		w.WriteHeader(http.StatusConflict)
	case code == codes.Unavailable:
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	default:
		MustSendError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(s.Message()))
	if err != nil {
		MustSendError(err, w)
	}
	return
}
