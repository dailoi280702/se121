package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/internal/recommender"
	"github.com/dailoi280702/se121/recommendation-service/internal/repo"
	"github.com/dailoi280702/se121/recommendation-service/pkg/recommendation"
)

type server struct {
	blogRepo    repo.BlogRepository
	tagRepo     repo.TagRepository
	recommender recommender.BlogRecommender
	recommendation.UnimplementedRecommendationServiceServer
}

func NewServer(blogRepo repo.BlogRepository, tagRepo repo.TagRepository) *server {
	return &server{
		blogRepo:    blogRepo,
		tagRepo:     tagRepo,
		recommender: recommender.BlogRecommender{},
	}
}

func serverError(err error) error {
	return fmt.Errorf("Blog server error: %v", err)
}

func (s *server) GetRelatedBlog(ctx context.Context, req *recommendation.GetRelatedBlogReq) (*recommendation.GetRelatedBlogRes, error) {
	s.recommender.DataSource = recommender.NewRelatedBlogsRecommender(s.blogRepo, s.tagRepo)

	id := strconv.Itoa(int(req.BlogId))
	recommendedBlogs, err := s.recommender.GetRecommendation(context.Background(), id, req.NumberOfBlog)
	if err != nil {
		return nil, serverError(err)
	}

	return &recommendation.GetRelatedBlogRes{Blogs: recommendedBlogs}, nil
}

func (s *server) GetUserRecommendedBlogs(ctx context.Context, req *recommendation.GetUserRecommendedBlogsReq) (*blog.Blogs, error) {
	s.recommender.DataSource = recommender.NewUserBlogsRecommender(s.blogRepo, s.tagRepo)

	recommendedBlogs, err := s.recommender.GetRecommendation(context.Background(), req.UserId, req.Limit)
	if err != nil {
		return nil, serverError(err)
	}

	return &blog.Blogs{Blogs: recommendedBlogs}, nil
}
