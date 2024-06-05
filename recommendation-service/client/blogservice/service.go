package blogservice

import (
	"context"
	"flag"
	"log"
	"sync"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var blogServicePort = flag.String("blogServicePort", "blog-service:50051", "the address to connect to blog service")

var (
	once     sync.Once
	instance *BlogService
)

type BlogService struct {
	blog.BlogServiceClient
	*grpc.ClientConn
}

func NewBlogService(ctx context.Context) (*grpc.ClientConn, blog.BlogServiceClient) {
	conn, err := grpc.DialContext(ctx, *blogServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect blog service: %v", err)
	}

	return conn, blog.NewBlogServiceClient(conn)
}

func GetInstance() *BlogService {
	once.Do(func() {
		conn, client := NewBlogService(context.Background())

		instance = &BlogService{
			BlogServiceClient: client,
			ClientConn:        conn,
		}
	})

	return instance
}
