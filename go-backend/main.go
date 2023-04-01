package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dailoi280702/se121/go_backend/api/router"
	"github.com/dailoi280702/se121/go_backend/protos"
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
	r := router.InitRouter(c)
	http.ListenAndServe(":8000", r)
}
