package main

import (
	"flag"
	"log"
	"net/http"

	api_v1 "github.com/dailoi280702/se121/go_backend/api/v1/router"
	"github.com/dailoi280702/se121/go_backend/protos"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "python-backend:50051", "the address to connect to")

func main() {
	// grpc
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewHelloClient(conn)

	// routes
	router := chi.NewRouter()
	router.Mount("/v1", api_v1.InitRouter(c))
	http.ListenAndServe(":8000", router)
}
