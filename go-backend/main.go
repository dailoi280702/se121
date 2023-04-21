package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	api_v1 "github.com/dailoi280702/se121/go_backend/api/v1/router"
	"github.com/dailoi280702/se121/go_backend/internal/service/user"
	"github.com/dailoi280702/se121/go_backend/protos"
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
	userServicePort = flag.String("userServicePort", "user-service:50051", "the address to connect to")
	redisAddr       = flag.String("redisAddr", "redis:6379", "the address to connect to redis")
)

func NewUserService(ctx context.Context) (*grpc.ClientConn, user.UserServiceClient) {
	conn, err := grpc.DialContext(ctx, *userServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect user service: %v", err)
	}

	return conn, user.NewUserServiceClient(conn)
}

func main() {
	// grpc
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect grpc: %v", err)
	}
	defer conn.Close()
	c := protos.NewHelloClient(conn)

	userServiceConn, userService := NewUserService(context.Background())
	defer userServiceConn.Close()

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
	defer db.Close()

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
	router.Mount("/v1", api_v1.InitRouter(c, redisClient, db, userService))
	log.Fatalf("Error serving api: %v", http.ListenAndServe(":8000", router))
}
