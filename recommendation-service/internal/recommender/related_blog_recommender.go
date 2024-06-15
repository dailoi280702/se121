package recommender

import (
	"context"
	"strconv"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/internal/repo"
)

type RelatedBlogsDataSource struct {
	blogRepo repo.BlogRepository
	tagRepo  repo.TagRepository
}

func (r *RelatedBlogsDataSource) GetRelatedBlogTags(ctx context.Context, idStr string, limit int32) ([]*blog.BlogTags, error) {
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return nil, err
	}

	return r.tagRepo.GetTagsFromBlogId(ctx, int32(id))
}

func (r *RelatedBlogsDataSource) GetLatestBlogTags(ctx context.Context) ([]*blog.BlogTags, error) {
	return r.tagRepo.GetLatestTags(ctx)
}

func (r *RelatedBlogsDataSource) GetBlogsFromIds(ctx context.Context, ids []int32) ([]*blog.Blog, error) {
	return r.blogRepo.GetBlogFromIds(ctx, ids)
}

func NewRelatedBlogsRecommender(blogRepo repo.BlogRepository, tagRepo repo.TagRepository) RecommenderDataSource {
	return &RelatedBlogsDataSource{
		blogRepo: blogRepo,
		tagRepo:  tagRepo,
	}
}
