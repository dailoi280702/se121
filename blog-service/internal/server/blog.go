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

// :TODO
func (s *server) GetBlog(ctx context.Context, req *blog.GetBlogReq) (*blog.Blog, error) {
	// Check for blog existence
	id := req.GetId()
	err := checkBlogExistence(s.db, id)
	if err != nil {
		return nil, err
	}

	// Fetch blog from database
	blog, err := getBlog(s.db, id)
	if err != nil {
		return nil, serverError(err)
	}

	return blog, nil
}

func (s *server) CreateBlog(ctx context.Context, req *blog.CreateBlogReq) (*blog.Empty, error) {
	// Validate and verify inputs
	err := validateBLog(s.db, &req.Title, &req.Body, req.Tldr, &req.Author, req.ImageUrl, req.Tags)
	if err != nil {
		return nil, err
	}

	// Create transaction
	tx, err := s.db.Begin()
	if err != nil {
		return nil, serverError(fmt.Errorf("failed to start transaction: %v", err))
	}

	// Insert blog into database
	if err := createBlogWithTags(tx, req); err != nil {
		if err != nil {
			log.Println("Insert blog failed, rooling back...")
			_ = tx.Rollback()
			return nil, serverError(err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println("Insert blog failed, rooling back...")
		_ = tx.Rollback()
		return nil, serverError(fmt.Errorf("failed to commit transaction: %v", err))
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
		tagIDs, err := insertTagsIfNotExists(tx, req.Tags)
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
		log.Println("Update blog failed, rooling back...")
		_ = tx.Rollback()
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
	if err := deleteBlog(s.db, int(id)); err != nil {
		return nil, serverError(err)
	}

	return &blog.Empty{}, nil
}

// :TODO
func (s *server) SearchForBlogs(ctx context.Context, req *blog.SearchReq) (*blog.SearchBlogsRes, error) {
	// Fetch list of blogs id and list of tags id
	blogIds, tagsIds, mapBlogTags, err := getBlogIdsAndTagIds(s.db, req)
	if err != nil {
		return nil, serverError(err)
	}

	var res blog.SearchBlogsRes
	var tags []*blog.Tag

	errCh := make(chan error)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// Fetach blogs
		total, err := getNumsOfBlogs(s.db, req)
		errCh <- err
		res.Total = int32(total)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		// Fetach tags
		res.Blogs, err = getBlogsByIds(s.db, blogIds)
		errCh <- err
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		tags, err = getTagsByIds(s.db, tagsIds)
		errCh <- err
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return nil, serverError(err)
		}
	}

	mapTags := map[int]*blog.Tag{}
	for _, t := range tags {
		mapTags[int((*t).Id)] = t
	}

	// Combines blogs and tags
	for _, b := range res.Blogs {
		tagsOfBlog, ok := mapBlogTags[int((*b).Id)]
		if ok {
			for _, id := range tagsOfBlog {
				b.Tags = append(b.Tags, mapTags[id])
			}
		}
	}

	return &res, nil
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

func validateBLog(db *sql.DB, title, body, tldr, author, imageUrl *string, tags []*blog.Tag) error {
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
	if tldr != nil {
		if strings.TrimSpace(*tldr) == "" {
			validationErrors["tldr"] = "TLDR can not be empty"
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
func createBlogWithTags(tx *sql.Tx, req *blog.CreateBlogReq) error {
	// Insert blog record
	blogId, err := insertBlog(tx, req)
	if err != nil {
		return err
	}
	if len(req.Tags) == 0 {
		return nil
	}

	tagIds, err := insertTagsIfNotExists(tx, req.Tags)
	if err != nil {
		return err
	}

	// Create blog and tags references
	err = createBlogTags(tx, blogId, tagIds)
	return err
}

// genreate sql query for searching blogs from grpc request as string
func generateSearchBlogQuery(sel string, req *blog.SearchReq) string {
	if sel == "" {
		panic("can not select nothing")
	}

	query := fmt.Sprintf(`
        SELECT %s 
		FROM blogs AS b
		LEFT JOIN blog_tags AS bt ON b.id = bt.blog_id
		LEFT JOIN tags AS t ON bt.tag_id = t.id
		`, sel)

	if req.GetQuery() != "" {
		query += fmt.Sprintf(`WHERE 1=1
            AND (b.title ILIKE '%%%s%%'
            OR b.body ILIKE '%%%s%%'
            OR t.name ILIKE '%%%s%%')`,
			req.GetQuery(), req.GetQuery(), req.GetQuery())
	}

	if req.GetOrderby() != "" {
		orderBy := "b.created_at"
		switch req.GetOrderby() {
		case "date":
			orderBy = "b.created_at"
		case "title":
			orderBy = "b.title"
		case "body":
			orderBy = "b.body"
		case "tldr":
			orderBy = "b.tldr"
		}
		query += fmt.Sprintf(" ORDER BY %s", orderBy)
		if req.GetIsAscending() {
			query += " ASC"
		} else {
			query += " DESC"
		}
	}

	if req.GetStartAt() > 0 {
		query += fmt.Sprintf(" OFFSET %d", req.GetStartAt())
	}

	if req.GetLimit() > 0 {
		query += fmt.Sprintf(" LIMIT %d", req.GetLimit())
	}

	return query
}

// return in order list of blog id, list of tag id and map of blog ids and its tags
func getBlogIdsAndTagIds(db *sql.DB, req *blog.SearchReq) ([]int, []int, map[int][]int, error) {
	blogIds := []int{}
	tagsIds := []int{}
	mapBlogTags := map[int][]int{}

	query := generateSearchBlogQuery("b.id, t.id", req)
	log.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return blogIds, tagsIds, mapBlogTags, nil
		}
		return nil, nil, nil, err
	}

	for rows.Next() {
		var bId int
		var tId *int
		if err := rows.Scan(&bId, &tId); err != nil {
			return nil, nil, nil, err
		}
		_, ok := mapBlogTags[bId]
		if !ok {
			mapBlogTags[bId] = []int{}
			blogIds = append(blogIds, bId)
		}
		if tId != nil {
			mapBlogTags[bId] = append(mapBlogTags[bId], *tId)
			tagsIds = append(tagsIds, *tId)
		}
	}

	return blogIds, removeDuplicates(tagsIds), mapBlogTags, nil
}

func getNumsOfBlogs(db *sql.DB, req *blog.SearchReq) (int, error) {
	query := generateSearchBlogQuery("COUNT(DISTINCT b.id)", &blog.SearchReq{
		Query:       req.Query,
		Orderby:     nil,
		IsAscending: nil,
		StartAt:     nil,
		Limit:       nil,
	})
	// query = fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS subquery", query)

	var count int
	if err := db.QueryRow(query).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}

func removeDuplicates(slice []int) []int {
	seen := make(map[int]struct{})
	result := make([]int, 0)

	for _, value := range slice {
		if _, ok := seen[value]; !ok {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}
