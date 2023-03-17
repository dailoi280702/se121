package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	// "time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/dailoi280702/se121/go_backend/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr    = flag.String("addr", "python-backend:50051", "the address to connect to")
	message = flag.String("name", "go", "Name to greet")
)

func main() {
	// grpc
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewHelloClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	// defer cancel()

	// router handler
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hello")); err != nil {
			log.Fatal(err)
		}
	})

	r.Get("/say-hello", func(w http.ResponseWriter, r *http.Request) {
		responce, err := c.SayHello(context.Background(), &protos.HelloRequest{Message: *message})
		if err != nil {
			msg := fmt.Sprintf("could not say hello: %v", err)
			if _, err := w.Write([]byte(msg)); err != nil {
				log.Fatal(err)
			}

			log.Panicf("could not say hello: %v", err)

			return
		}

		msg := fmt.Sprintf("Greeting: %s", responce.GetMessage())
		if _, err := w.Write([]byte(msg)); err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":8000", r)
}
