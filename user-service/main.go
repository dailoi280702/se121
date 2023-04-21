package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/dailoi280702/se121/user_service/internal/service"
	"github.com/dailoi280702/se121/user_service/userpb"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/status"
	// "github.com/golang/protobuf/proto"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

type userServer struct {
	service *service.Service
	userpb.UnimplementedUserServiceServer
}

func (s *userServer) GetUser(c context.Context, req *userpb.GetUserReq) (*userpb.GetUserRes, error) {
	user, err := s.service.GetUser(req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Service err while creatting user %s", err.Error()))
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &userpb.GetUserRes{User: &userpb.User{
		Id:       user.Id,
		Name:     user.Name,
		Email:    &user.Email,
		ImageUrl: &user.ImageUrl,
		CreateAt: user.CreateAt.Unix(),
		IsAdmin:  user.IsAdmin,
	}}, nil
}

func (s *userServer) VerifyUser(context.Context, *userpb.VerifyUserReq) (*userpb.VerifyUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUser not implemented")
}

func (s *userServer) GetUsers(*userpb.GetUsersReq, userpb.UserService_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}

func (s *userServer) CreateUser(context.Context, *userpb.CreateUserReq) (*userpb.CreateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (s *userServer) UpdateUser(context.Context, *userpb.User) (*userpb.UpdateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func newServer(service *service.Service) *userServer {
	s := &userServer{service: service}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = data.Path("x509/server_cert.pem")
		}
		if *keyFile == "" {
			*keyFile = data.Path("x509/server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()
	userpb.RegisterUserServiceServer(grpcServer, newServer(&service.Service{DB: db}))
	grpcServer.Serve(lis)
}
