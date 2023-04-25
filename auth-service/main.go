package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"

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

func (s *AuthServiceServer) SignIn(ctx context.Context, req *auth.SignInReq) (*auth.User, error) {
	details := make(map[string]string)
	nameOrEmail := req.GetNameOrEmail()
	password := req.GetPassword()

	// validate input
	if nameOrEmail == "" {
		details["nameOrEmail"] = "user name or email cannot be empty"
	} else {
		emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		usernameRegex := regexp.MustCompile("^[A-Za-z0-9]+(?:[ _-][A-Za-z0-9]+)*$")
		isEmail := emailRegex.MatchString(nameOrEmail)
		isUsername := usernameRegex.MatchString(nameOrEmail)
		if !isEmail && !isUsername {
			details["nameOrEmail"] = "neither user name is password are valid"
		}
	}
	if password == "" {
		details["password"] = "password cannot be empty"
	}

	if len(details) != 0 {
		data, err := json.Marshal(details)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("auth service err: %v", err))
		}
		return nil, status.Error(codes.InvalidArgument, string(data))
	}

	// // verify user
	// var data *user.User
	// if len(details) == 0 {
	// 	data, err := s.userService.VerifyUser(context.Background(), &user.VerifyUserReq{
	// 		NameOrEmail: nameOrEmail,
	// 		Passord:     password,
	// 	})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

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
