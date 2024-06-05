package recommender

import (
	"context"
	"strconv"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/internal/dao"
)

type RelatedBlogsRecommender struct {
	DefaultBlogRecommender
}

func (r *RelatedBlogsRecommender) GetInteractedBlogTags(ctx context.Context, idStr string, limit int32) ([]*blog.BlogTags, error) {
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return nil, err
	}

	return r.tagDao.GetTagsFromBlogId(ctx, int32(id))
}

func NewRelatedBlogsRecommender(blogDao dao.BlogDao, tagDao dao.TagDao) BlogRecommender {
	defaultRecommender := &DefaultBlogRecommender{
		blogDao: blogDao,
		tagDao:  tagDao,
	}

	recommender := &RelatedBlogsRecommender{
		DefaultBlogRecommender: *defaultRecommender,
	}

	defaultRecommender.BlogRecommender = recommender

	return defaultRecommender
}
