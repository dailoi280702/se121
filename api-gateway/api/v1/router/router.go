package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dailoi280702/se121/api-gateway/api/v1/handlers"
	"github.com/dailoi280702/se121/api-gateway/internal/service/auth"
	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/dailoi280702/se121/comment-service/pkg/comment"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/text_generate"
	"github.com/dailoi280702/se121/recommendation-service/pkg/recommendation"
	"github.com/dailoi280702/se121/search-service/pkg/search"
	"github.com/dailoi280702/se121/user-service/userpb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/redis/go-redis/v9"
)

func InitRouter(
	redisClient *redis.Client,
	db *sql.DB,
	userService user.UserServiceClient,
	authService auth.AuthServiceClient,
	carService car.CarServiceClient,
	blogService blog.BlogServiceClient,
	commentService comment.CommentServiceClient,
	searchService search.SearchServiceClient,
	textGenerateService text_generate.TextGenerateServiceClient,
	recommendationService recommendation.RecommendationServiceClient,
) *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
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

	// router.Mount("/say-hello", handlers.NewHelloRouter(gprcHelloClient).Routes())
	router.Mount("/auth", handlers.NewAuthHandler(redisClient, db, authService).Routes())
	router.Mount("/car", handlers.NewCarRoutes(carService))
	router.Mount("/brand", handlers.NewBrandRoutes(carService))
	router.Mount("/series", handlers.NewSeriesRoutes(carService))
	router.Mount("/blog", handlers.NewBlogRoutes(blogService, recommendationService))
	router.Mount("/tag", handlers.NewTagRoutes(blogService))
	router.Mount("/comment", handlers.NewCommentRoutes(commentService))
	router.Mount("/search", handlers.NewSearchRoutes(searchService))
	router.Mount("/text-generate", handlers.NewTextGenerateRoutes(textGenerateService))
	router.Mount("/user", handlers.NewUserHanlder(userService))

	return router
}
