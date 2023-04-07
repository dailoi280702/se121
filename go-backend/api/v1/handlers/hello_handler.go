package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dailoi280702/se121/go_backend/protos"
	"github.com/go-chi/chi/v5"
)

type HelloHandler struct {
	grpcHelloClient protos.HelloClient
}

func NewHelloRouter(grpcHelloClient protos.HelloClient) *HelloHandler {
	return &HelloHandler{
		grpcHelloClient: grpcHelloClient,
	}
}

func (h HelloHandler) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", h.SayHelloFromPython)

	return router
}

func (h HelloHandler) SayHello(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Peekapoo")); err != nil {
		log.Panicf("something's wrong %v", err)
	}
}

func (h HelloHandler) SayHelloFromPython(w http.ResponseWriter, r *http.Request) {
	response, err := h.grpcHelloClient.SayHello(context.Background(), &protos.HelloRequest{Message: "go"})
	if err != nil {
		msg := fmt.Sprintf("could not say hello: %v", err)
		if _, err := w.Write([]byte(msg)); err != nil {
			log.Fatal(err)
		}

		log.Panicf("could not say hello: %v", err)

		return
	}

	msg := fmt.Sprintf("Peekapoo: %s", response.GetMessage())
	if _, err := w.Write([]byte(msg)); err != nil {
		log.Fatal(err)
	}
}
