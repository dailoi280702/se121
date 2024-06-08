package recommender

import (
	"context"
	"strconv"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/internal/repo"
)

type RelatedBlogsRecommender struct {
	DefaultBlogRecommender
}

func (r *RelatedBlogsRecommender) GetInteractedBlogTags(ctx context.Context, idStr string, limit int32) ([]*blog.BlogTags, error) {
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return nil, err
	}

	return r.tagRepo.GetTagsFromBlogId(ctx, int32(id))
}

func NewRelatedBlogsRecommender(blogRepo repo.BlogRepository, tagRepo repo.TagRepository) BlogRecommender {
	defaultRecommender := &DefaultBlogRecommender{
		blogRepo: blogRepo,
		tagRepo:  tagRepo,
	}

	recommender := &RelatedBlogsRecommender{
		DefaultBlogRecommender: *defaultRecommender,
	}

	defaultRecommender.BlogRecommender = recommender

	return defaultRecommender
}
