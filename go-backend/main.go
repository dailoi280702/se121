package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	api_v1 "github.com/dailoi280702/se121/go_backend/api/v1/router"
	"github.com/dailoi280702/se121/go_backend/protos"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "python-backend:50051", "the address to connect to")

func main() {
	// grpc
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewHelloClient(conn)

	// redis
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// send SET command
	err = client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	// send GET command and print the value
	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)

	// routes
	router := chi.NewRouter()
	router.Mount("/v1", api_v1.InitRouter(c))
	http.ListenAndServe(":8000", router)
}
