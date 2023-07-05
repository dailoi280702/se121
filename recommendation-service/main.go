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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddress   = flag.String("server address", "[::]:50051", "address of recommendation server")
	blogServicePort = flag.String("blogServicePort", "blog-service:50051", "the address to connect to blog service")
)

func NewBlogService(ctx context.Context) (*grpc.ClientConn, blog.BlogServiceClient) {
	conn, err := grpc.DialContext(ctx, *blogServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect blog service: %v", err)
	}

	return conn, blog.NewBlogServiceClient(conn)
}

func serveServer(db *sql.DB, blogService blog.BlogServiceClient) {
	lis, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sv := grpc.NewServer()
	recommendation.RegisterRecommendationServiceServer(sv, server.NewServer(db, blogService))
	if err := sv.Serve(lis); err != nil {
		log.Fatalf("failed to serve server: %v", err)
	}
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to db :%v", err)
	}

	blogServiceConn, blogService := NewBlogService(context.Background())
	defer func() {
		db.Close()
		blogServiceConn.Close()
	}()

	serveServer(db, blogService)
}
