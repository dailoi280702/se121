package server

import (
	"database/sql"
	"fmt"
	"strings"

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

type errorResponse struct {
	Messages []string          `json:"messages,omitempty"`
	Details  map[string]string `json:"details,omitempty"`
}

func insertTagIfNotExists(db *sql.DB, tag *blog.Tag) (int, error) {
	if tag == nil {
		return 0, nil
	}

	// Clean the tag name by removing extra spaces and convert to lowercase
	cleanTagName := strings.ToLower(strings.TrimSpace(tag.Name))

	// Check if the tag already exists (case-insensitive and space-insensitive)
	query := "SELECT id FROM tags WHERE LOWER(TRIM(name)) = $1"
	var id int
	err := db.QueryRow(query, cleanTagName).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			// If the tag does not exist, insert a new record
			insertQuery := "INSERT INTO tags (name, description) VALUES ($1, $2) RETURNING id"
			err = db.QueryRow(insertQuery, tag.Name, tag.Description).Scan(&id)
			if err != nil {
				return 0, fmt.Errorf("failed to insert tag: %v", err)
			}
		} else {
			return 0, fmt.Errorf("failed to insert tag: %v", err)
		}
	}

	return id, nil
}

func createBlogTags(db *sql.DB, blogId int, tags []int) error {
	// Prepare the INSERT statement
	insertQuery := "INSERT INTO blog_tags (tag_id, blog_id) VALUES "

	// Build the value placeholders and argument list
	values := []interface{}{}
	placeholders := []string{}
	for _, tagId := range tags {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", len(values)+1, len(values)+2))
		values = append(values, tagId, blogId)
	}

	// Combine the query and placeholders
	insertQuery += strings.Join(placeholders, ", ")

	// Execute the INSERT statement
	_, err := db.Exec(insertQuery, values...)
	if err != nil {
		return fmt.Errorf("failed to insert into blog_tags: %v", err)
	}

	return nil
}

func insertBlog(db *sql.DB, req *blog.CreateBlogReq) (int, error) {
	var id int

	// Execute the INSERT statement and retrieve the ID of the newly inserted blog
	err := db.QueryRow(`
		INSERT INTO blogs (title, body, author, tldr, image_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, req.Title, req.Body, req.Author, req.Tldr, req.ImageUrl).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert blog to database: %v", err)
	}

	return id, nil
}
