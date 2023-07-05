package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"github.com/dailoi280702/se121/user-service/internal/service"
	"github.com/dailoi280702/se121/user-service/userpb"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/status"
	// "github.com/golang/protobuf/proto"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 50051, "The server port")
)

type userServer struct {
	service *service.Service
	user.UnimplementedUserServiceServer
}

func (s *userServer) GetUser(c context.Context, req *user.GetUserReq) (*user.GetUserRes, error) {
	u, err := s.service.GetUser(req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Service err while getting user %s", err.Error()))
	}

	if u == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &user.GetUserRes{User: &user.User{
		Id:       u.Id,
		Name:     u.Name,
		Email:    u.Email,
		ImageUrl: u.ImageUrl,
		CreateAt: u.CreateAt.UnixMilli(),
		IsAdmin:  u.IsAdmin,
	}}, nil
}

func (s *userServer) GetUserProfilesByIds(c context.Context, req *user.GetUserProfilesByIdsReq) (*user.GetUserProfilesByIdsRes, error) {
	userProfiles, err := s.service.GetUserProfilesByIds(req.Ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "user service error while fetching users by ids: %v", err)
	}
	return &user.GetUserProfilesByIdsRes{Users: userProfiles}, nil
}

func (s *userServer) VerifyUser(ctx context.Context, req *user.VerifyUserReq) (*user.User, error) {
	name := strings.TrimSpace(req.GetNameOrEmail())
	password := strings.TrimSpace(req.GetPassord())

	if name == "" || password == "" {
		return nil, status.Error(codes.InvalidArgument, "inputs cannot be empty")
	}

	u, err := s.service.VerifyUser(name, password)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrIncorrectNameEmailOrPassword):
			return nil, status.Error(codes.NotFound, "user name, email or password is not correct")
		default:
			return nil, status.Errorf(codes.Internal, "user service error while verifying user: %v", err)
		}
	}

	return &user.User{
		Id:       u.Id,
		Name:     u.Name,
		Email:    u.Email,
		ImageUrl: u.ImageUrl,
		CreateAt: u.CreateAt.UnixMilli(),
		IsAdmin:  u.IsAdmin,
	}, nil
}

func (s *userServer) GetUsers(*user.GetUsersReq, user.UserService_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}

func (s *userServer) CreateUser(req *user.CreateUserReq, stream user.UserService_CreateUserServer) error {
	name, email, password := req.GetName(), req.GetEmail(), req.GetPassword()
	err := s.service.AddUser(service.User{
		Name:     name,
		Email:    &email,
		Password: password,
	})
	if err != nil {
		var validationErrors *service.ValidationErrors

		switch {
		case errors.As(err, &validationErrors):
			if len(validationErrors.Messages) > 0 {
				if err := stream.Send(&user.CreateUserRes{Errors: validationErrors.Messages}); err != nil {
					return status.Error(codes.Internal, "user service error: cannot send data while creatting user")
				}
				return status.Error(codes.AlreadyExists, "user existed")
			}
		default:
			return status.Error(codes.Internal, fmt.Sprintf("Service err while creatting user %s", err.Error()))
		}
	}

	return nil
}

func (s *userServer) UpdateUser(context.Context, *user.User) (*user.UpdateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func (s *userServer) MarkBlogAsReaded(ctx context.Context, req *user.MarkBlogAsReadedReq) (*utils.Empty, error) {
	// Check if the user_id and blog_id combination already exists
	existsQuery := `
		SELECT EXISTS(
			SELECT 1 FROM readed_blogs WHERE user_id = $1 AND blog_id = $2
		)
	`

	var exists bool
	err := s.service.DB.QueryRowContext(ctx, existsQuery, req.UserId, req.BlogId).Scan(&exists)
	if err != nil {
		return nil, err
	}

	if exists {
		// Update the existing row with the current time
		updateQuery := `
			UPDATE readed_blogs SET at = now() WHERE user_id = $1 AND blog_id = $2
		`

		_, err = s.service.DB.ExecContext(ctx, updateQuery, req.UserId, req.BlogId)
		if err != nil {
			return nil, err
		}
	} else {
		// Insert a new row with the current time
		insertQuery := `
			INSERT INTO readed_blogs (user_id, blog_id, at) VALUES ($1, $2, now())
		`

		_, err = s.service.DB.ExecContext(ctx, insertQuery, req.UserId, req.BlogId)
		if err != nil {
			return nil, err
		}
	}

	return &utils.Empty{}, nil
}

func (s *userServer) GetRecentlyReadedBlogsIds(ctx context.Context, req *user.GetRecentlyReadedBlogsIdsReq) (*user.GetRecentlyReadedBlogsIdsRes, error) {
	// Prepare the SQL statement to select the most recently read blog IDs for the given user
	query := `
		SELECT blog_id
		FROM readed_blogs
		WHERE user_id = $1
		ORDER BY at DESC
		LIMIT $2
	`

	// Execute the SQL statement
	rows, err := s.service.DB.QueryContext(ctx, query, req.UserId, req.Limit)
	if err != nil {
		if err == sql.ErrNoRows {
			defer rows.Close()
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	// Collect the blog IDs from the query result
	var blogIds []int32
	for rows.Next() {
		var blogID int32
		err := rows.Scan(&blogID)
		if err != nil {
			return nil, err
		}
		blogIds = append(blogIds, blogID)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Create the response message with the retrieved blog IDs
	res := &user.GetRecentlyReadedBlogsIdsRes{
		BlogIds: blogIds,
	}

	return res, nil
}

func newServer(service *service.Service) *userServer {
	s := &userServer{service: service}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("[::]:%d", *port))
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
	user.RegisterUserServiceServer(grpcServer, newServer(service.NewService(db)))
	log.Fatalf("Error serving user service: %v", grpcServer.Serve(lis))
}
