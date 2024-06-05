package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/internal/dao"
	"github.com/dailoi280702/se121/recommendation-service/internal/recommender"
	"github.com/dailoi280702/se121/recommendation-service/pkg/recommendation"
)

type server struct {
	blogDao dao.BlogDao
	tagDao  dao.TagDao
	recommendation.UnimplementedRecommendationServiceServer
}

func NewServer(blogDao dao.BlogDao, tagDao dao.TagDao) *server {
	return &server{
		blogDao: blogDao,
		tagDao:  tagDao,
	}
}

func serverError(err error) error {
	return fmt.Errorf("blog server error: %v", err)
}

func (s *server) GetRelatedBlog(ctx context.Context, req *recommendation.GetRelatedBlogReq) (*recommendation.GetRelatedBlogRes, error) {
	recommender := recommender.NewRelatedBlogsRecommender(s.blogDao, s.tagDao)

	id := strconv.Itoa(int(req.BlogId))
	recommendedBlogs, err := recommender.GetRecommendation(context.Background(), id, req.NumberOfBlog)
	if err != nil {
		return nil, serverError(err)
	}

	return &recommendation.GetRelatedBlogRes{
		Blogs: recommendedBlogs,
	}, nil
}

func (s *server) GetUserRecommendedBlogs(ctx context.Context, req *recommendation.GetUserRecommendedBlogsReq) (*blog.Blogs, error) {
	recommender := recommender.NewUserBlogsRecommender(s.blogDao, s.tagDao)

	recommendedBlogs, err := recommender.GetRecommendation(context.Background(), req.UserId, req.Limit)
	if err != nil {
		return nil, serverError(err)
	}

	return &blog.Blogs{Blogs: recommendedBlogs}, nil
}