package server

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"strings"
	"sync"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/pkg/go/sqlutils"
)

type server struct {
	db *sql.DB
	blog.UnimplementedBlogServiceServer
}

func NewServer(db *sql.DB) *server {
	return &server{db: db}
}

func serverError(err error) error {
	return fmt.Errorf("Blog server error: %v", err)
}

type errorResponse struct {
	Messages []string          `json:"messages,omitempty"`
	Details  map[string]string `json:"details,omitempty"`
}

func insertTagIfNotExists(tx *sql.Tx, tag *blog.Tag) (int, error) {
	if tag == nil {
		return 0, nil
	}
	log.Println("connected")

	// Clean the tag name by removing extra spaces and convert to lowercase
	cleanTagName := strings.ToLower(strings.TrimSpace(tag.Name))

	// Check if the tag already exists (case-insensitive and space-insensitive)
	query := "SELECT id FROM tags WHERE LOWER(TRIM(name)) = $1"
	var id int
	err := tx.QueryRow(query, cleanTagName).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			// If the tag does not exist, insert a new record
			insertQuery := "INSERT INTO tags (name, description) VALUES ($1, $2) RETURNING id"
			err = tx.QueryRow(insertQuery, tag.Name, tag.Description).Scan(&id)
			if err != nil {
				return 0, fmt.Errorf("failed to insert tag: %v", err)
			}
		} else {
			return 0, fmt.Errorf("failed to insert tag: %v", err)
		}
	}

	return id, nil
}

func createBlogTags(tx *sql.Tx, blogId int, tags []int) error {
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
	_, err := tx.Exec(insertQuery, values...)
	if err != nil {
		return fmt.Errorf("failed to insert into blog_tags: %v", err)
	}

	return nil
}

func insertBlog(tx *sql.Tx, req *blog.CreateBlogReq) (int, error) {
	var id int

	// Execute the INSERT statement and retrieve the ID of the newly inserted blog
	err := tx.QueryRow(`
		INSERT INTO blogs (title, body, author, tldr, image_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, req.Title, req.Body, req.Author, req.Tldr, req.ImageUrl).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert blog to database: %v", err)
	}

	return id, nil
}

func updateBlog(tx *sql.Tx, req *blog.UpdateBlogReq) error {
	// Prepare update data
	updateData := map[string]interface{}{"updated_at": "NOW()"}
	if req.Title != nil {
		updateData["title"] = *req.Title
	}
	if req.Body != nil {
		updateData["body"] = *req.Body
	}
	if req.ImageUrl != nil {
		updateData["image_url"] = *req.ImageUrl
	}
	if req.Tldr != nil {
		updateData["tldr"] = *req.Tldr
	}

	// Update blog record
	return sqlutils.UpdateRecordWithTransaction(tx, "blogs", updateData, int(req.Id))
}

func updateBlogTags(tx *sql.Tx, blogID int, tagIDs []int) error {
	// Remove old blog_tags
	_, err := tx.Exec("DELETE FROM blog_tags WHERE blog_id = $1", blogID)
	if err != nil {
		return err
	}

	if len(tagIDs) == 0 {
		return nil
	}

	// Prepare the INSERT statement for multiple blog_tags
	insertStmt := "INSERT INTO blog_tags (tag_id, blog_id) VALUES "
	values := make([]string, len(tagIDs))
	for i := range tagIDs {
		values[i] = fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2)
	}

	insertQuery := insertStmt + strings.Join(values, ", ")
	args := make([]interface{}, len(tagIDs)*2)
	for i, tagID := range tagIDs {
		args[i*2] = tagID
		args[i*2+1] = blogID
	}

	// Insert new blog_tags
	_, err = tx.Exec(insertQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func deleteBlog(db *sql.DB, blogID int) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if p := recover(); p != nil {
			// Rollback the transaction in case of a panic
			_ = tx.Rollback()
			panic(p) // Re-throw panic after rollback
		} else if err != nil {
			// Rollback the transaction in case of an error
			_ = tx.Rollback()
		} else {
			// Commit the transaction if everything succeeded
			err = tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				log.Println("failed to commit transaction:", err)
			}
		}
	}()

	// Delete associated records in blog_cars
	_, err = tx.Exec("DELETE FROM blog_cars WHERE blog_id = $1", blogID)
	if err != nil {
		return fmt.Errorf("failed to delete blog_cars records: %v", err)
	}

	// Delete associated records in blog_tags
	_, err = tx.Exec("DELETE FROM blog_tags WHERE blog_id = $1", blogID)
	if err != nil {
		return fmt.Errorf("failed to delete blog_tags records: %v", err)
	}

	// Delete associated records in blog_comments
	_, err = tx.Exec("DELETE FROM blog_comments WHERE blog_id = $1", blogID)
	if err != nil {
		return fmt.Errorf("failed to delete blog_comments records: %v", err)
	}

	// Delete the blog record itself
	_, err = tx.Exec("DELETE FROM blogs WHERE id = $1", blogID)
	if err != nil {
		return fmt.Errorf("failed to delete blog record: %v", err)
	}

	return nil
}

// Insert new tag record if tag name not exists and return list of tags id
func insertTagsIfNotExists(tx *sql.Tx, tags []*blog.Tag) ([]int, error) {
	// numsWorker := getNumWorkers(len(tags))
	numsWorker := 1
	jobs := make(chan *blog.Tag, len(tags))
	errCh := make(chan error)
	var wg sync.WaitGroup
	wg.Add(numsWorker)

	tagIds := []int{}

	// Function to insert tags concurrently
	worker := func() {
		defer wg.Done()
		for tag := range jobs {
			tagId, err := insertTagIfNotExists(tx, tag)
			errCh <- err
			tagIds = append(tagIds, tagId)
		}
	}

	// Spawn workers in goroutines
	for i := 0; i < numsWorker; i++ {
		go worker()
	}

	// Send jobs to workers
	for _, tag := range tags {
		jobs <- tag
	}
	close(jobs)

	// Close errCh after all workers fished working
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Check for errors in the errCh channel
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return tagIds, nil
}

// calculates the number of worker goroutines based on the number of jobs.
func getNumWorkers(numJobs int) int {
	numWorkers := int(math.Floor(math.Sqrt(float64(numJobs))))
	if numWorkers > 10 {
		return 10
	}
	return numWorkers
}
