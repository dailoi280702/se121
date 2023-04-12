package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dailoi280702/se121/go_backend/api/v1/handlers"
	"github.com/dailoi280702/se121/go_backend/protos"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/redis/go-redis/v9"
)

func InitRouter(gprcHelloClient protos.HelloClient, redisClient *redis.Client, db *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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
	router.Mount("/auth", handlers.NewAuthHandler(redisClient, db).Routes())

	return router
}
