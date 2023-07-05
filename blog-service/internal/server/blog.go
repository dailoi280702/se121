package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/pkg/go/grpc/generated/utils"
	"github.com/dailoi280702/se121/pkg/go/sqlutils"
	u "github.com/dailoi280702/se121/pkg/go/utils"
	"github.com/lib/pq"
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

func (s *server) CreateBlog(ctx context.Context, req *blog.CreateBlogReq) (*blog.CreateBlogRes, error) {
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
	id, err := createBlogWithTags(tx, req)
	if err != nil {
		log.Printf("Insert blog failed, rooling back... : %v", err)
		_ = tx.Rollback()
		return nil, serverError(err)
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Insert blog failed, rooling back... : %v", err)
		_ = tx.Rollback()
		return nil, serverError(fmt.Errorf("failed to commit transaction: %v", err))
	}
	log.Println("New blog inserted")

	return &blog.CreateBlogRes{Id: int32(id)}, nil
}

func (s *server) UpdateBlog(ctx context.Context, req *blog.UpdateBlogReq) (*utils.Empty, error) {
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

	return &utils.Empty{}, nil
}

func (s *server) DeleteBlog(ctx context.Context, req *blog.DeleteBlogReq) (*utils.Empty, error) {
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

	return &utils.Empty{}, nil
}

// :TODO
func (s *server) SearchForBlogs(ctx context.Context, req *utils.SearchReq) (*blog.SearchBlogsRes, error) {
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

func (s *server) GetNumberOfBlogs(context.Context, *utils.Empty) (*blog.GetNumberOfBlogsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberOfBlogs not implemented")
}

func (s *server) GetBlogsFromIds(ctx context.Context, req *blog.BlogIds) (*blog.Blogs, error) {
	ids := []int{}
	for _, id := range req.Ids {
		ids = append(ids, int(id))
	}
	blogs, err := getBlogsByIds(s.db, ids)
	if err != nil {
		return nil, serverError(err)
	}
	return &blog.Blogs{Blogs: blogs}, nil
}

func (s *server) GetTagsFromBlogIds(context context.Context, req *blog.BlogIds) (*blog.GetTagsFromBlogIdsRes, error) {
	// Initialize the response struct
	response := &blog.GetTagsFromBlogIdsRes{
		BlogTags: []*blog.BlogTags{},
	}

	blogTagMap, err := getBlogTagsFromBlogIds(s.db, req.Ids)
	if err != nil {
		return nil, serverError(err)
	}

	// Build the final response using the blogTagMap
	for _, id := range req.Ids {
		tags := blogTagMap[id]
		response.BlogTags = append(response.BlogTags, &blog.BlogTags{
			BlogId: id,
			Tags:   tags,
		})
	}

	return response, nil
}

func (s *server) GetLatestBlogTags(context context.Context, req *blog.GetLatestBlogTagsReq) (*blog.GetLatestBlogTagsRes, error) {
	// Initialize the response struct
	response := &blog.GetLatestBlogTagsRes{
		BlogTags: []*blog.BlogTags{},
	}

	ids, err := getLatestBlogIDs(s.db, req.GetNumberOfBlogs)
	if err != nil {
		return nil, serverError(err)
	}

	blogTagMap, err := getBlogTagsFromBlogIds(s.db, ids)
	if err != nil {
		return nil, serverError(err)
	}

	// Build the final response using the blogTagMap
	for _, id := range ids {
		tags := blogTagMap[id]
		response.BlogTags = append(response.BlogTags, &blog.BlogTags{
			BlogId: id,
			Tags:   tags,
		})
	}

	return response, nil
}

func checkBlogExistence(db *sql.DB, id int32) error {
	exists, err := sqlutils.IdExists(db, "blogs", id)
	if err != nil {
		return serverError(err)
	}
	if !exists {
		return u.ConvertGrpcToJsonError(codes.NotFound, fmt.Sprintf("Blog %d does not exist", id))
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
		return u.ConvertGrpcToJsonError(codes.InvalidArgument, errorResponse{
			Details: validationErrors,
		})
	}

	return nil
}

// creates a blog record with associated tags in the database.
func createBlogWithTags(tx *sql.Tx, req *blog.CreateBlogReq) (int, error) {
	// Insert blog record
	blogId, err := insertBlog(tx, req)
	if err != nil {
		return 0, err
	}
	if len(req.Tags) == 0 {
		return blogId, nil
	}

	tagIds, err := insertTagsIfNotExists(tx, req.Tags)
	if err != nil {
		return blogId, err
	}

	// Create blog and tags references
	err = createBlogTags(tx, blogId, tagIds)
	return blogId, err
}

// genreate sql query for searching blogs from grpc request as string
func generateSearchBlogQuery(sel string, req *utils.SearchReq) string {
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

	return query
}

// return in order list of blog id, list of tag id and map of blog ids and its tags
func getBlogIdsAndTagIds(db *sql.DB, req *utils.SearchReq) ([]int, []int, map[int][]int, error) {
	blogIds := []int{}
	tagsIds := []int{}
	mapBlogTags := map[int][]int{}

	query := generateSearchBlogQuery("b.id, t.id", req)
	rows, err := db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return blogIds, tagsIds, mapBlogTags, nil
		}
		return nil, nil, nil, err
	}

	tempBlogs := map[int]struct{}{}
	startAt := int(req.GetStartAt())
	limit := int(req.GetLimit())

	for rows.Next() {
		var bId int
		var tId *int
		if err := rows.Scan(&bId, &tId); err != nil {
			return nil, nil, nil, err
		}
		if limit != 0 && len(mapBlogTags) >= limit {
			break
		}
		if len(tempBlogs)+1 < startAt {
			tempBlogs[bId] = struct{}{}
			continue
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

func getNumsOfBlogs(db *sql.DB, req *utils.SearchReq) (int, error) {
	query := generateSearchBlogQuery("COUNT(DISTINCT b.id)", &utils.SearchReq{
		Query:       req.Query,
		Orderby:     nil,
		IsAscending: nil,
		StartAt:     nil,
		Limit:       nil,
	})

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

func getLatestBlogIDs(db *sql.DB, n int32) ([]int32, error) {
	// Prepare the SQL statement to fetch the IDs of the N latest blogs
	query := `
		SELECT id
		FROM blogs
		ORDER BY created_at DESC
		LIMIT $1
	`

	// Execute the SQL statement to retrieve the IDs of the N latest blogs
	rows, err := db.Query(query, n)
	if err != nil {
		if err == sql.ErrNoRows {
			return []int32{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	// Initialize the list of blog IDs
	ids := []int32{}

	// Iterate over the result rows
	for rows.Next() {
		var blogID int32

		// Scan the blog ID from the row into a variable
		err := rows.Scan(&blogID)
		if err != nil {
			return nil, err
		}

		// Append the blog ID to the list
		ids = append(ids, blogID)
	}

	// Check for any errors occurred during rows iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

func getBlogTagsFromBlogIds(db *sql.DB, blogIds []int32) (map[int32][]*blog.Tag, error) {
	blogTagMap := make(map[int32][]*blog.Tag)

	// Prepare the SQL statement to fetch blog tags by blog IDs
	query := `
		SELECT blog_tags.blog_id, tags.id, tags.name, tags.description
		FROM blog_tags
		INNER JOIN tags ON blog_tags.tag_id = tags.id
		WHERE blog_id = ANY($1)
	`

	// Execute the SQL statement with the provided blog IDs
	rows, err := db.Query(query, pq.Array(blogIds))
	if err != nil {
		if err == sql.ErrNoRows {
			return blogTagMap, nil
		}
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result rows
	for rows.Next() {
		var blogID int32
		var tag blog.Tag

		// Scan the values from the row into variables
		err := rows.Scan(&blogID, &tag.Id, &tag.Name, &tag.Description)
		if err != nil {
			return nil, err
		}

		// Append the tag to the corresponding blog ID entry in the map
		blogTagMap[blogID] = append(blogTagMap[blogID], &tag)
	}

	// Check for any errors occurred during rows iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return blogTagMap, nil
}
