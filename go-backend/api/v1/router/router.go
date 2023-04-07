package router

import (
	"log"
	"net/http"

	"github.com/dailoi280702/se121/go_backend/api/v1/handlers"
	"github.com/dailoi280702/se121/go_backend/protos"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

func InitRouter(gprcHelloClient protos.HelloClient, redisClient *redis.Client) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.CleanPath)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hello")); err != nil {
			log.Fatal(err)
		}
	})

	router.Mount("/say-hello", handlers.NewHelloRouter(gprcHelloClient).Routes())
	router.Mount("/auth", handlers.NewAuthHandler(redisClient).Routes())

	return router
}
