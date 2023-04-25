package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/dailoi280702/se121/auth_service/internal/service"
	"github.com/dailoi280702/se121/auth_service/internal/service/auth"
	"github.com/dailoi280702/se121/auth_service/internal/service/user"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	serverPort         = flag.Int("server port", 50051, "The server port")
	serverNetwork      = flag.String("server network", "[::]:", "The server network")
	userServicePort    = flag.Int("user service port", 50051, "The user service port")
	userServiceNetwork = flag.String("user service network", "user-service", "The user service network")
	redisAddr          = flag.String("redisAddr", "redis:6379", "the address to connect to redis")
)

type AuthServiceServer struct {
	service     *service.Service
	userService user.UserServiceClient

	auth.UnimplementedAuthServiceServer
}

func (s *AuthServiceServer) SignIn(context.Context, *auth.SignInReq) (*auth.SignInRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}

func (s *AuthServiceServer) SignUp(context.Context, *auth.SignUpReq) (*auth.SignUpRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}

func (s *AuthServiceServer) Refresh(context.Context, *auth.RefreshReq) (*auth.RefreshRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

func (s *AuthServiceServer) SignOut(context.Context, *auth.SignOutReq) (*auth.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}

func NewUserService(ctx context.Context) (*grpc.ClientConn, user.UserServiceClient) {
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s%d", *userServiceNetwork, *userServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect user service: %v", err)
	}

	return conn, user.NewUserServiceClient(conn)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s%d", *serverNetwork, *serverPort))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	userServiceConn, userService := NewUserService(context.Background())
	defer userServiceConn.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     *redisAddr,
		Password: "",
		DB:       0,
	})
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, &AuthServiceServer{
		service:     service.NewService(redisClient),
		userService: userService,
	})
	log.Fatalf("Error serving auth service: %v", grpcServer.Serve(lis))
}
