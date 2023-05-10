package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/dailoi280702/se121/car-service/internal/server"
	"github.com/dailoi280702/se121/car-service/pkg/car"
	"google.golang.org/grpc"
)

var serverAdress = flag.String("server address", "[::]:5050", "server address of car serivce")

func main() {
	lis, err := net.Listen("tcp", *serverAdress)
	if err != nil {
		fmt.Printf("failed to listen on %s: %v", *serverAdress, err)
	}
	sv := grpc.NewServer()
	car.RegisterCarServiceServer(sv, server.NewServer())
	if err = sv.Serve(lis); err != nil {
		fmt.Printf("failed to serve car server: %v", err)
	}
}
