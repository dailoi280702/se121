package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net"
	"os"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/internal/server"
	"github.com/dailoi280702/se121/recommendation-service/pkg/recommendation"
	"github.com/dailoi280702/se121/user-service/userpb"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddress   = flag.String("server address", "[::]:50051", "address of recommendation server")
	blogServicePort = flag.String("blogServicePort", "blog-service:50051", "the address to connect to blog service")
	userServicePort = flag.String("userServicePort", "user-service:50051", "the address to connect to user service")
)

func NewUserService(ctx context.Context) (*grpc.ClientConn, user.UserServiceClient) {
	conn, err := grpc.DialContext(ctx, *userServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect user service: %v", err)
	}

	return conn, user.NewUserServiceClient(conn)
}

func NewBlogService(ctx context.Context) (*grpc.ClientConn, blog.BlogServiceClient) {
	conn, err := grpc.DialContext(ctx, *blogServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect blog service: %v", err)
	}

	return conn, blog.NewBlogServiceClient(conn)
}

func serveServer(db *sql.DB, blogService blog.BlogServiceClient, userService user.UserServiceClient) {
	lis, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sv := grpc.NewServer()
	recommendation.RegisterRecommendationServiceServer(sv, server.NewServer(db, blogService, userService))
	if err := sv.Serve(lis); err != nil {
		log.Fatalf("failed to serve server: %v", err)
	}
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to db :%v", err)
	}

	ctx := context.Background()
	blogServiceConn, blogService := NewBlogService(ctx)
	userServiceConn, userService := NewUserService(ctx)
	defer func() {
		db.Close()
		userServiceConn.Close()
		blogServiceConn.Close()
	}()

	serveServer(db, blogService, userService)
}
