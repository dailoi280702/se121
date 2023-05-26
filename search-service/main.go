package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"github.com/dailoi280702/se121/search-service/pkg/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAdress    = flag.String("server address", "[::]:50051", "server address of car serivce")
	carServicePort  = flag.String("carServicePort", "car-service:50051", "the address to connect to car service")
	blogServicePort = flag.String("blogServicePort", "blog-service:50051", "the address to connect to blog service")
)

type server struct {
	carService  car.CarServiceClient
	blogService blog.BlogServiceClient
	search.UnimplementedSearchServiceServer
}

func newServer(carService car.CarServiceClient, blogService blog.BlogServiceClient) *server {
	return &server{
		carService:  carService,
		blogService: blogService,
	}
}

func (s *server) Search(ctx context.Context, req *utils.SearchReq) (*search.SearchRes, error) {
	res := search.SearchRes{}

	errCh := make(chan error)
	var err error
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		res.Blogs, err = s.blogService.SearchForBlogs(ctx, req)
		errCh <- err
		wg.Done()
	}()

	go func() {
		res.Cars, err = s.carService.SearchForCar(ctx, req)
		errCh <- err
		wg.Done()
	}()

	go func() {
		res.Brands, err = s.carService.SearchForBrand(ctx, req)
		errCh <- err
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			log.Panicf("Search service got and err: %v", err)
		}
	}

	return &res, nil
}

func main() {
	ctx := context.Background()
	carServiceConn, carService := NewCarService(ctx)
	blogServiceConn, blogService := NewBlogService(ctx)

	defer func() {
		carServiceConn.Close()
		blogServiceConn.Close()
	}()

	serveGrpcServer(carService, blogService)
}

func serveGrpcServer(carService car.CarServiceClient, blogService blog.BlogServiceClient) {
	lis, err := net.Listen("tcp", *serverAdress)
	if err != nil {
		fmt.Printf("failed to listen on %s: %v", *serverAdress, err)
	}
	sv := grpc.NewServer()
	search.RegisterSearchServiceServer(sv, newServer(carService, blogService))
	if err = sv.Serve(lis); err != nil {
		fmt.Printf("failed to serve car server: %v", err)
	}
}

func NewCarService(ctx context.Context) (*grpc.ClientConn, car.CarServiceClient) {
	conn, err := grpc.DialContext(ctx, *carServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect car service: %v", err)
	}

	return conn, car.NewCarServiceClient(conn)
}

func NewBlogService(ctx context.Context) (*grpc.ClientConn, blog.BlogServiceClient) {
	conn, err := grpc.DialContext(ctx, *blogServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect blog service: %v", err)
	}

	return conn, blog.NewBlogServiceClient(conn)
}
