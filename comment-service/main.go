package main

import (
	"database/sql"
	"flag"
	"log"
	"net"
	"os"

	"github.com/dailoi280702/se121/comment-service/internal/server"
	"github.com/dailoi280702/se121/comment-service/pkg/comment"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var serverAddress = flag.String("server address", "[::]:50051", "address of comment server")

func main() {
	lis, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("failed to connect to databse: %v", err)
	}
	defer db.Close()

	sv := grpc.NewServer()
	comment.RegisterCommentServiceServer(sv, server.NewServer(db))
	println("Serving comment service")
	if err := sv.Serve(lis); err != nil {
		log.Fatalf("failed to serve server: %v", err)
	}
}
