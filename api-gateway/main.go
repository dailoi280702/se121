package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	api_v1 "github.com/dailoi280702/se121/api-gateway/api/v1/router"
	"github.com/dailoi280702/se121/api-gateway/internal/service/auth"
	"github.com/dailoi280702/se121/api-gateway/internal/service/user"
	"github.com/dailoi280702/se121/api-gateway/protos"
	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr            = flag.String("addr", "python-backend:50051", "the address to connect to")
	userServicePort = flag.String("userServicePort", "user-service:50051", "the address to connect to user service")
	authServicePort = flag.String("authServicePort", "auth-service:50051", "the address to connect to auth service")
	carServicePort  = flag.String("carServicePort", "car-service:50051", "the address to connect to car service")
	blogServicePort = flag.String("blogServicePort", "blog-service:50051", "the address to connect to blog service")
	redisAddr       = flag.String("redisAddr", "redis:6379", "the address to connect to redis")
)

func NewUserService(ctx context.Context) (*grpc.ClientConn, user.UserServiceClient) {
	conn, err := grpc.DialContext(ctx, *userServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect user service: %v", err)
	}

	return conn, user.NewUserServiceClient(conn)
}

func NewAuthService(ctx context.Context) (*grpc.ClientConn, auth.AuthServiceClient) {
	conn, err := grpc.DialContext(ctx, *authServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect auth service: %v", err)
	}

	return conn, auth.NewAuthServiceClient(conn)
}

func NewCarService(ctx context.Context) (*grpc.ClientConn, car.CarServiceClient) {
	conn, err := grpc.DialContext(ctx, *carServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect auth service: %v", err)
	}

	return conn, car.NewCarServiceClient(conn)
}

func NewBlogService(ctx context.Context) (*grpc.ClientConn, blog.BlogServiceClient) {
	conn, err := grpc.DialContext(ctx, *blogServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect auth service: %v", err)
	}

	return conn, blog.NewBlogServiceClient(conn)
}

func main() {
	// grpc
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect grpc: %v", err)
	}
	c := protos.NewHelloClient(conn)

	ctx := context.Background()

	userServiceConn, userService := NewUserService(ctx)
	authServiceConn, authService := NewAuthService(ctx)
	carServiceConn, carService := NewCarService(ctx)
	blogServiceConn, blogService := NewBlogService(ctx)

	// redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     *redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	defer func() {
		db.Close()
		conn.Close()
		userServiceConn.Close()
		authServiceConn.Close()
		carServiceConn.Close()
		blogServiceConn.Close()
	}()

	// database migratetion
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("can not init migration: %v", err)
	}

	if err := m.Up(); err != nil {
		switch {
		case errors.Is(err, migrate.ErrNoChange):
			log.Println("migration: no change")
		default:
			log.Fatalf("failed to migrate db up: %v", err)
		}
	}

	// // create table if it does not exist
	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// routes
	router := chi.NewRouter()
	router.Mount("/v1", api_v1.InitRouter(
		c, redisClient, db,
		userService,
		authService,
		carService,
		blogService))
	log.Fatalf("Error serving api: %v", http.ListenAndServe(":8000", router))
}
