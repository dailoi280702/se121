package recommender

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/dailoi280702/se121/blog-service/pkg/blog"
	"github.com/dailoi280702/se121/recommendation-service/internal/repo"
)

type BlogRecommender interface {
	GetInteractedBlogTags(ctx context.Context, id string, limit int32) ([]*blog.BlogTags, error)
	GetLatestBlogTags(ctx context.Context) ([]*blog.BlogTags, error)
	GetBlogsFromIds(ctx context.Context, ids []int32) ([]*blog.Blog, error)
	GetRecommendation(ctx context.Context, id string, limit int32) ([]*blog.Blog, error)
}

var _ BlogRecommender = (*DefaultBlogRecommender)(nil)

type DefaultBlogRecommender struct {
	blogRepo repo.BlogRepository
	tagRepo  repo.TagRepository
	BlogRecommender
}

func (r *DefaultBlogRecommender) GetLatestBlogTags(ctx context.Context) ([]*blog.BlogTags, error) {
	return r.tagRepo.GetLatestTags(ctx)
}

func (r *DefaultBlogRecommender) GetBlogsFromIds(ctx context.Context, ids []int32) ([]*blog.Blog, error) {
	return r.blogRepo.GetBlogFromIds(ctx, ids)
}

func (r *DefaultBlogRecommender) GetRecommendation(ctx context.Context, id string, limit int32) ([]*blog.Blog, error) {
	sourceItems := []Item{}
	destItems := []Item{}

	errCh := make(chan error)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		res, err := r.GetInteractedBlogTags(context.Background(), id, limit)
		errCh <- err
		for _, value := range res {
			tags := []int32{}
			for _, tag := range value.Tags {
				tags = append(tags, tag.Id)
			}
			sourceItems = append(sourceItems, Item{ID: int(value.BlogId), Tags: tags})
		}
		wg.Done()
	}()

	go func() {
		res, err := r.GetLatestBlogTags(context.Background())
		errCh <- err
		for _, value := range res {
			tags := []int32{}
			for _, tag := range value.Tags {
				tags = append(tags, tag.Id)
			}
			destItems = append(destItems, Item{ID: int(value.BlogId), Tags: tags})
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return nil, fmt.Errorf("Recommendation service error: %v", err)
		}
	}

	recommendations := generateRecommendations(sourceItems, destItems, int(limit))

	recommendedBlogIds := []int32{}
	for _, r := range recommendations {
		recommendedBlogIds = append(recommendedBlogIds, int32(r.ItemID))
	}

	blogs, err := r.GetBlogsFromIds(ctx, recommendedBlogIds)
	if err != nil {
		return nil, fmt.Errorf("Recommendation service error while fetch recommendedBlogs: %v", err)
	}

	return blogs, nil
}

type Item struct {
	ID   int
	Tags []int32
}

type Recommendation struct {
	ItemID     int
	Similarity float64
}

type ByTags []Item

func (a ByTags) Len() int           { return len(a) }
func (a ByTags) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTags) Less(i, j int) bool { return len(a[i].Tags) > len(a[j].Tags) }

func calculateSimilarity(item Item, other Item) float64 {
	// Calculate similarity between clickedItem and item based on their tags
	// This can be your custom similarity calculation logic
	// Here, we assume a simple similarity score based on the number of common tags
	commonTags := 0
	for _, tag := range item.Tags {
		for _, itemTag := range other.Tags {
			if tag == itemTag {
				commonTags++
				break
			}
		}
	}
	totalTags := len(item.Tags) + len(other.Tags) - commonTags
	return float64(commonTags) / float64(totalTags)
}

func generateRecommendations(clickedItems []Item, items []Item, numRecommendations int) []Recommendation {
	recommendations := make([]Recommendation, 0)

	// Calculate similarity for the last clicked items
	for _, clickedItem := range clickedItems {
		for _, item := range items {
			similarity := calculateSimilarity(clickedItem, item)
			recommendation := Recommendation{
				ItemID:     item.ID,
				Similarity: similarity,
			}
			recommendations = append(recommendations, recommendation)
		}
	}

	// Sort recommendations based on similarity
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Similarity > recommendations[j].Similarity
	})

	// Keep only the top N recommendations
	if len(recommendations) > numRecommendations {
		recommendations = recommendations[:numRecommendations]
	}

	return recommendations
}
