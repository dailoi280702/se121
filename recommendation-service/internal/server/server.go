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
	blogRepo repo.BlogRepository
	tagRepo  repo.TagRepository
	recommendation.UnimplementedRecommendationServiceServer
}

func NewServer(blogRepo repo.BlogRepository, tagRepo repo.TagRepository) *server {
	return &server{
		blogRepo: blogRepo,
		tagRepo:  tagRepo,
	}
}

func serverError(err error) error {
	return fmt.Errorf("Blog server error: %v", err)
}

func (s *server) GetRelatedBlog(ctx context.Context, req *recommendation.GetRelatedBlogReq) (*recommendation.GetRelatedBlogRes, error) {
	recommender := recommender.NewRelatedBlogsRecommender(s.blogRepo, s.tagRepo)

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
	recommender := recommender.NewUserBlogsRecommender(s.blogRepo, s.tagRepo)

	recommendedBlogs, err := recommender.GetRecommendation(context.Background(), req.UserId, req.Limit)
	if err != nil {
		return nil, serverError(err)
	}

	return &blog.Blogs{Blogs: recommendedBlogs}, nil
}

// func (s *server) GetRelatedBlog(ctx context.Context, req *recommendation.GetRelatedBlogReq) (*recommendation.GetRelatedBlogRes, error) {
// 	log.Println("Get related blog request recived")
// 	sourceItems := []Item{}
// 	destItems := []Item{}
//
// 	errCh := make(chan error)
// 	var wg sync.WaitGroup
// 	wg.Add(2)
//
// 	go func() {
// 		res, err := s.tagrepo.GetTagsFromBlogId(context.Background(), req.BlogId)
// 		errCh <- err
// 		for _, value := range res {
// 			tags := []int32{}
// 			for _, tag := range value.Tags {
// 				tags = append(tags, tag.Id)
// 			}
// 			sourceItems = append(sourceItems, Item{ID: int(value.BlogId), Tags: tags})
// 		}
// 		wg.Done()
// 	}()
//
// 	go func() {
// 		res, err := s.tagrepo.GetLatestTags(context.Background())
// 		errCh <- err
// 		for _, value := range res {
// 			tags := []int32{}
// 			for _, tag := range value.Tags {
// 				tags = append(tags, tag.Id)
// 			}
// 			destItems = append(destItems, Item{ID: int(value.BlogId), Tags: tags})
// 		}
// 		wg.Done()
// 	}()
//
// 	go func() {
// 		wg.Wait()
// 		close(errCh)
// 	}()
//
// 	for err := range errCh {
// 		if err != nil {
// 			return nil, serverError(fmt.Errorf("Recommendation service error: %v", err))
// 		}
// 	}
//
// 	recommendations := generateRecommendations(sourceItems, destItems, int(req.NumberOfBlog))
//
// 	recommendedBlogIds := []int32{}
// 	for _, r := range recommendations {
// 		recommendedBlogIds = append(recommendedBlogIds, int32(r.ItemID))
// 	}
//
// 	blogs, err := s.blogrepo.GetBlogFromIds(ctx, recommendedBlogIds)
// 	if err != nil {
// 		return nil, serverError(fmt.Errorf("Recommendation service error while fetch recommendedBlogs: %v", err))
// 	}
//
// 	return &recommendation.GetRelatedBlogRes{Blogs: blogs}, nil
// }
//
// func (s *server) GetUserRecommendedBlogs(ctx context.Context, req *recommendation.GetUserRecommendedBlogsReq) (*blog.Blogs, error) {
// 	sourceItems := []Item{}
// 	destItems := []Item{}
//
// 	errCh := make(chan error)
// 	var wg sync.WaitGroup
// 	wg.Add(2)
//
// 	go func() {
// 		res, err := s.tagrepo.GetUserTagsFromRecentActivity(ctx, req.GetUserId(), req.GetLimit())
// 		errCh <- err
// 		for _, value := range res {
// 			tags := []int32{}
// 			for _, tag := range value.Tags {
// 				tags = append(tags, tag.Id)
// 			}
// 			sourceItems = append(sourceItems, Item{ID: int(value.BlogId), Tags: tags})
// 		}
// 		wg.Done()
// 	}()
//
// 	go func() {
// 		res, err := s.tagrepo.GetLatestTags(ctx)
// 		errCh <- err
// 		for _, value := range res {
// 			tags := []int32{}
// 			for _, tag := range value.Tags {
// 				tags = append(tags, tag.Id)
// 			}
// 			destItems = append(destItems, Item{ID: int(value.BlogId), Tags: tags})
// 		}
// 		wg.Done()
// 	}()
//
// 	go func() {
// 		wg.Wait()
// 		close(errCh)
// 	}()
//
// 	for err := range errCh {
// 		if err != nil {
// 			return nil, serverError(fmt.Errorf("Recommendation service error: %v", err))
// 		}
// 	}
//
// 	recommendations := generateRecommendations(sourceItems, destItems, int(req.Limit))
//
// 	recommendedBlogIds := []int32{}
// 	for _, r := range recommendations {
// 		recommendedBlogIds = append(recommendedBlogIds, int32(r.ItemID))
// 	}
//
// 	blogs, err := s.blogrepo.GetBlogFromIds(ctx, recommendedBlogIds)
// 	if err != nil {
// 		return nil, serverError(fmt.Errorf("Recommendation service error while fetch recommendedBlogs: %v", err))
// 	}
//
// 	return &blog.Blogs{Blogs: blogs}, nil
// }
//
// type Item struct {
// 	ID   int
// 	Tags []int32
// }
//
// type Recommendation struct {
// 	ItemID     int
// 	Similarity float64
// }
//
// type ByTags []Item
//
// func (a ByTags) Len() int           { return len(a) }
// func (a ByTags) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByTags) Less(i, j int) bool { return len(a[i].Tags) > len(a[j].Tags) }
//
// func calculateSimilarity(item Item, other Item) float64 {
// 	// Calculate similarity between clickedItem and item based on their tags
// 	// This can be your custom similarity calculation logic
// 	// Here, we assume a simple similarity score based on the number of common tags
// 	commonTags := 0
// 	for _, tag := range item.Tags {
// 		for _, itemTag := range other.Tags {
// 			if tag == itemTag {
// 				commonTags++
// 				break
// 			}
// 		}
// 	}
// 	totalTags := len(item.Tags) + len(other.Tags) - commonTags
// 	return float64(commonTags) / float64(totalTags)
// }
//
// func generateRecommendations(clickedItems []Item, items []Item, numRecommendations int) []Recommendation {
// 	recommendations := make([]Recommendation, 0)
//
// 	// Calculate similarity for the last clicked items
// 	for _, clickedItem := range clickedItems {
// 		for _, item := range items {
// 			similarity := calculateSimilarity(clickedItem, item)
// 			recommendation := Recommendation{
// 				ItemID:     item.ID,
// 				Similarity: similarity,
// 			}
// 			recommendations = append(recommendations, recommendation)
// 		}
// 	}
//
// 	// Sort recommendations based on similarity
// 	sort.Slice(recommendations, func(i, j int) bool {
// 		return recommendations[i].Similarity > recommendations[j].Similarity
// 	})
//
// 	// Keep only the top N recommendations
// 	if len(recommendations) > numRecommendations {
// 		recommendations = recommendations[:numRecommendations]
// 	}
//
// 	return recommendations
// }
