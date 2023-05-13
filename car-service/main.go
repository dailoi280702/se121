package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/dailoi280702/se121/car-service/internal/server"
	"github.com/dailoi280702/se121/car-service/pkg/car"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var serverAdress = flag.String("server address", "[::]:5050", "server address of car serivce")

func serveGrpcServer(db *sql.DB) {
	lis, err := net.Listen("tcp", *serverAdress)
	if err != nil {
		fmt.Printf("failed to listen on %s: %v", *serverAdress, err)
	}
	sv := grpc.NewServer()
	car.RegisterCarServiceServer(sv, server.NewServer(db))
	if err = sv.Serve(lis); err != nil {
		fmt.Printf("failed to serve car server: %v", err)
	}
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	defer func() {
		db.Close()
	}()

	serveGrpcServer(db)
}
