package userservice

import (
	"context"
	"flag"
	"log"
	"sync"

	"github.com/dailoi280702/se121/user-service/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var userServicePort = flag.String("userServicePort", "user-service:50051", "the address to connect to user service")

var (
	once     sync.Once
	instance *UserService
)

type UserService struct {
	user.UserServiceClient
	*grpc.ClientConn
}

func NewUserService(ctx context.Context) (*grpc.ClientConn, user.UserServiceClient) {
	conn, err := grpc.DialContext(ctx, *userServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect user service: %v", err)
	}

	return conn, user.NewUserServiceClient(conn)
}

func GetInstance() *UserService {
	once.Do(func() {
		conn, client := NewUserService(context.Background())

		instance = &UserService{
			UserServiceClient: client,
			ClientConn:        conn,
		}
	})

	return instance
}
