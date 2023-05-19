package server

import (
	"database/sql"
	"fmt"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
)

type server struct {
	db *sql.DB
	blog.UnimplementedBlogServiceServer
}

func NewServer(db *sql.DB) *server {
	return &server{db: db}
}

func serverError(err error) error {
	return fmt.Errorf("Blog server error %v", err)
}
