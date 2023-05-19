package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/pkg/go/sqlutils"
	"github.com/dailoi280702/se121/pkg/go/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetBlog(ctx context.Context, req *blog.GetBlogReq) (*blog.Blog, error) {
	// Check for blog existence
	id := req.GetId()
	err := checkBlogExistence(s.db, id)
	if err != nil {
		return nil, err
	}

	// Fetch blog from database

	return nil, status.Errorf(codes.Unimplemented, "method GetBlog not implemented")
}

func (s *server) CreateBlog(ctx context.Context, req *blog.CreateBlogReq) (*blog.Empty, error) {
	// Validate and verify inputs
	err := validateBLog(s.db, &req.Title, &req.Body, req.Tldr, &req.Author, req.ImageUrl, req.Tags)
	if err != nil {
		return nil, err
	}

	// Insert blog into database
	if err := createBlogWithTags(s.db, req); err != nil {
		return nil, serverError(err)
	}

	return &blog.Empty{}, nil
}

func (s *server) UpdateBlog(ctx context.Context, req *blog.UpdateBlogReq) (*blog.Empty, error) {
	// Check for blog existence
	id := req.GetId()
	err := checkBlogExistence(s.db, id)
	if err != nil {
		return nil, err
	}

	// Validate and verify inputs
	err = validateBLog(s.db, req.Title, req.Body, req.Tldr, nil, req.ImageUrl, req.Tags)
	if err != nil {
		return nil, err
	}

	errCh := make(chan error)
	var wg sync.WaitGroup
	wg.Add(2)

	tx, err := s.db.Begin()
	if err != nil {
		return nil, serverError(fmt.Errorf("failed to start transaction: %v", err))
	}

	go func() {
		err := updateBlog(tx, req)
		errCh <- err
		wg.Done()
	}()

	go func() {
		tagIDs, err := insertTagsIfNotExists(s.db, req.Tags)
		errCh <- err
		errCh <- updateBlogTags(tx, int(id), tagIDs)
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			log.Println("Update failed, rooling back...")
			_ = tx.Rollback()
			return nil, serverError(err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, serverError(fmt.Errorf("failed to commit transaction: %v", err))
	}

	return &blog.Empty{}, nil
}

func (s *server) DeleteBlog(ctx context.Context, req *blog.DeleteBlogReq) (*blog.Empty, error) {
	// Check for blog existence
	id := req.GetId()
	err := checkBlogExistence(s.db, id)
	if err != nil {
		return nil, err
	}

	// Delete blog record

	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}

func (s *server) SearchForBlogs(ctx context.Context, req *blog.SearchReq) (*blog.SearchBlogsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForBlogs not implemented")
}

func (s *server) GetNumberOfBlogs(context.Context, *blog.Empty) (*blog.GetNumberOfBlogsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberOfBlogs not implemented")
}

func checkBlogExistence(db *sql.DB, id int32) error {
	exists, err := sqlutils.IdExists(db, "blogs", id)
	if err != nil {
		return serverError(err)
	}
	if !exists {
		return utils.ConvertGrpcToJsonError(codes.NotFound, fmt.Sprintf("Blog %d does not exist", id))
	}
	return nil
}

func validateBLog(db *sql.DB, title, body, trdl, author, imageUrl *string, tags []*blog.Tag) error {
	validationErrors := map[string]string{}

	if title != nil {
		if strings.TrimSpace(*title) == "" {
			validationErrors["title"] = "Title can not be empty"
		}
	}
	if body != nil {
		if strings.TrimSpace(*body) == "" {
			validationErrors["body"] = "Body can not be empty"
		}
	}
	if trdl != nil {
		if strings.TrimSpace(*trdl) == "" {
			validationErrors["trdl"] = "TRDL can not be empty"
		}
	}
	if author != nil {
		if strings.TrimSpace(*author) == "" {
			validationErrors["author"] = "Author can not be empty"
		}
	}
	if imageUrl != nil {
		if strings.TrimSpace(*imageUrl) == "" {
			validationErrors["imageUrl"] = "Image URL can not be empty"
		}
	}

	if len(validationErrors) > 0 {
		return utils.ConvertGrpcToJsonError(codes.InvalidArgument, errorResponse{
			Details: validationErrors,
		})
	}

	return nil
}

// creates a blog record with associated tags in the database.
func createBlogWithTags(db *sql.DB, req *blog.CreateBlogReq) error {
	// Insert blog record
	blogId, err := insertBlog(db, req)
	if err != nil {
		return err
	}
	if len(req.Tags) == 0 {
		return nil
	}

	tagIds, err := insertTagsIfNotExists(db, req.Tags)
	if err != nil {
		return err
	}

	// Create blog and tags references
	err = createBlogTags(db, blogId, tagIds)
	return err
}
