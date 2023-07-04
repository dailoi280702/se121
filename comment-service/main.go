package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net"
	"os"

	"github.com/dailoi280702/se121/comment-service/internal/server"
	"github.com/dailoi280702/se121/comment-service/pkg/comment"
	user "github.com/dailoi280702/se121/user-service/userpb"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddress   = flag.String("server address", "[::]:50051", "address of comment server")
	userServicePort = flag.String("userServicePort", "user-service:50051", "the address to connect to user service")
)

func NewUserService(ctx context.Context) (*grpc.ClientConn, user.UserServiceClient) {
	conn, err := grpc.DialContext(ctx, *userServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect user service: %v", err)
	}

	return conn, user.NewUserServiceClient(conn)
}

func main() {
	lis, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to databse: %v", err)
	}
	defer db.Close()

	userServiceConn, userService := NewUserService(context.Background())
	defer userServiceConn.Close()

	sv := grpc.NewServer()
	comment.RegisterCommentServiceServer(sv, server.NewServer(db, userService))
	println("Serving comment service")
	if err := sv.Serve(lis); err != nil {
		log.Fatalf("failed to serve server: %v", err)
	}
}
