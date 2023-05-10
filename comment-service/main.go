package main

import (
	"flag"
	"log"
	"net"

	"github.com/dailoi280702/se121/comment-service/internal/server"
	"github.com/dailoi280702/se121/comment-service/pkg/comment"
	"google.golang.org/grpc"
)

var serverAddress = flag.String("server address", "[::]:50051", "address of comment server")

func main() {
	lis, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sv := grpc.NewServer()
	comment.RegisterCommentServiceServer(sv, server.NewServer())
	println("Serving comment service")
	if err := sv.Serve(lis); err != nil {
		log.Fatalf("failed to serve server: %v", err)
	}
}
