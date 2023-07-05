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
	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/car-service/pkg/car"
	"github.com/dailoi280702/se121/comment-service/pkg/comment"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/text_generate"
	"github.com/dailoi280702/se121/recommendation-service/pkg/recommendation"
	"github.com/dailoi280702/se121/search-service/pkg/search"
	"github.com/dailoi280702/se121/user-service/userpb"
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
	textGenerateServicePort   = flag.String("textGenerateServicePort", "text-generate-service:50051", "the address to connect to text generate service")
	userServicePort           = flag.String("userServicePort", "user-service:50051", "the address to connect to user service")
	authServicePort           = flag.String("authServicePort", "auth-service:50051", "the address to connect to auth service")
	carServicePort            = flag.String("carServicePort", "car-service:50051", "the address to connect to car service")
	blogServicePort           = flag.String("blogServicePort", "blog-service:50051", "the address to connect to blog service")
	commentServicePort        = flag.String("commentServicePort", "comment-service:50051", "the address to connect to comment service")
	searchServicePort         = flag.String("searchServicePort", "search-service:50051", "the address to connect to search service")
	recommendationServicePort = flag.String("recommendationServicePort", "recoomendation-service:50051", "the address to connect to recommendation service")
	redisAddr                 = flag.String("redisAddr", "redis:6379", "the address to connect to redis")
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
		log.Fatalf("failed to connect car service: %v", err)
	}

	return conn, car.NewCarServiceClient(conn)
}

func NewBlogService(ctx context.Context) (*grpc.ClientConn, blog.BlogServiceClient) {
	conn, err := grpc.DialContext(ctx, *blogServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect blog service: %v", err)
	}

	return conn, blog.NewBlogServiceClient(conn)
}

func NewTextGenerateService(ctx context.Context) (*grpc.ClientConn, text_generate.TextGenerateServiceClient) {
	conn, err := grpc.DialContext(ctx, *textGenerateServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect text generate service: %v", err)
	}

	return conn, text_generate.NewTextGenerateServiceClient(conn)
}

func NewCommentService(ctx context.Context) (*grpc.ClientConn, comment.CommentServiceClient) {
	conn, err := grpc.DialContext(ctx, *commentServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect comment service: %v", err)
	}

	return conn, comment.NewCommentServiceClient(conn)
}

func NewSearchService(ctx context.Context) (*grpc.ClientConn, search.SearchServiceClient) {
	conn, err := grpc.DialContext(ctx, *searchServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect search service: %v", err)
	}

	return conn, search.NewSearchServiceClient(conn)
}

func NewRecommendationService(ctx context.Context) (*grpc.ClientConn, recommendation.RecommendationServiceClient) {
	conn, err := grpc.DialContext(ctx, *recommendationServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect search service: %v", err)
	}

	return conn, recommendation.NewRecommendationServiceClient(conn)
}

func main() {
	// grpc
	ctx := context.Background()

	userServiceConn, userService := NewUserService(ctx)
	authServiceConn, authService := NewAuthService(ctx)
	carServiceConn, carService := NewCarService(ctx)
	blogServiceConn, blogService := NewBlogService(ctx)
	commentServiceConn, commentService := NewCommentService(ctx)
	searchServiceConn, searchService := NewSearchService(ctx)
	textGenerateServiceConn, textGenerateService := NewTextGenerateService(ctx)
	recommendationServiceConn, recommendationService := NewRecommendationService(ctx)

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
		userServiceConn.Close()
		authServiceConn.Close()
		carServiceConn.Close()
		blogServiceConn.Close()
		commentServiceConn.Close()
		searchServiceConn.Close()
		textGenerateServiceConn.Close()
		recommendationServiceConn.Close()
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

	// routes
	router := chi.NewRouter()
	router.Mount("/v1", api_v1.InitRouter(
		redisClient, db,
		userService,
		authService,
		carService,
		blogService,
		commentService,
		searchService,
		textGenerateService,
		recommendationService,
	))
	log.Fatalf("Error serving api: %v", http.ListenAndServe(":8000", router))
}
