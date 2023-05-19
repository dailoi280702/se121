package main

import (
	"database/sql"
	"errors"
	"flag"
	"log"
	"net"
	"os"

	"github.com/dailoi280702/se121/blog-service/internal/server"
	"github.com/dailoi280702/se121/blog-service/pkg/blog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var serverAddress = flag.String("server address", "[::]:50051", "address of blog server")

func serveServer(db *sql.DB) {
	lis, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sv := grpc.NewServer()
	blog.RegisterBlogServiceServer(sv, server.NewServer(db))
	if err := sv.Serve(lis); err != nil {
		log.Fatalf("failed to serve server: %v", err)
	}
}

func runDBMigration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create db diver :%v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver)
	if err != nil {
		log.Fatalf("failed init migration: %v", err)
	}
	err = m.Up()
	if err != nil {
		switch {
		case errors.Is(err, migrate.ErrNoChange):
			log.Println("migration: no change")
		default:
			log.Fatalf("failed to migrate db up: %v", err)
		}
	}
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to db :%v", err)
	}
	defer db.Close()

	runDBMigration(db)
	serveServer(db)
}
