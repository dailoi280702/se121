package main

import (
	"flag"
	"log"
	"net"

	"github.com/dailoi280702/se121/blog-service/internal/server"
	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	// "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"google.golang.org/grpc"
)

var serverAddress = flag.String("server address", "[::]:50051", "address of blog server")

func serveServer() {
	lis, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sv := grpc.NewServer()
	blog.RegisterBlogServiceServer(sv, server.NewServer())
	if err := sv.Serve(lis); err != nil {
		log.Fatalf("failed to serve server: %v", err)
	}
}

func main() {
	serveServer()
}
