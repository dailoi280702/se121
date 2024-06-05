package main

import (
	"flag"
	"log"
	"net"

	"github.com/dailoi280702/se121/recommendation-service/cache"
	"github.com/dailoi280702/se121/recommendation-service/client/blogservice"
	"github.com/dailoi280702/se121/recommendation-service/client/userservice"
	"github.com/dailoi280702/se121/recommendation-service/internal/dao"
	proxyserver "github.com/dailoi280702/se121/recommendation-service/internal/proxy_server"
	"github.com/dailoi280702/se121/recommendation-service/internal/server"
	"github.com/dailoi280702/se121/recommendation-service/pkg/recommendation"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var serverAddress = flag.String("server address", "[::]:50051", "address of recommendation server")

func main() {
	lis, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	blogService := blogservice.GetInstance()
	userService := userservice.GetInstance()
	cache := cache.GetInstance()

	defer func() {
		blogService.Close()
		userService.Close()
	}()

	blogRepo := dao.NewBlogRepository()
	tagRepo := dao.NewTagRepository()

	grpcServer := grpc.NewServer()
	server := server.NewServer(blogRepo, tagRepo)
	proxyServer := proxyserver.NewProxyServer(server, cache)

	recommendation.RegisterRecommendationServiceServer(grpcServer, proxyServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve server: %v", err)
	}
}
