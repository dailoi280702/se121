package cached_store

import (
	"context"
	"encoding/json"
	"flag"
	"time"

	"github.com/dailoi280702/se121/go_backend/models"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var (
	existingTokens = flag.String("existing tokens key", "existingTokens", "key for hashset of existing auth tokens")
	expiredTokens  = flag.String("expired tokens key", "expiredTokens", "key for hashset of expired auth tokens")
)

// An implementation of TokenStore which its data are stored in machine memory
type InMemoryTokenStore struct {
	client *redis.Client
}

func NewRedisAuthTokenStore(client *redis.Client) *InMemoryTokenStore {
	return &InMemoryTokenStore{
		client: client,
	}
}

func (s *InMemoryTokenStore) NewToken(userId string, isAdmin bool, lifetime time.Duration) (string, error) {
	key := uuid.NewString()
	token := &models.AuthToken{
		Token:     key,
		UserId:    userId,
		Admin:     isAdmin,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(lifetime),
	}
	err := s.client.HSet(context.Background(), *existingTokens, key, token).Err()
	return key, err
}

func (s *InMemoryTokenStore) IsExisting(token string) (bool, error) {
	return s.client.HExists(context.Background(), *existingTokens, token).Result()
}

func (s *InMemoryTokenStore) IsExpired(token string) (bool, error) {
	existing, err := s.IsExisting(token)
	if err != nil {
		return false, err
	}

	if existing {
		authToken, err := getToken(s.client, *existingTokens, token)
		if err != nil {
			return false, err
		}

		expired := authToken.ExpiresAt.Before(time.Now())
		if expired {
			if err := s.Remove(token); err != nil {
				return false, err
			}
			return true, nil
		}

		return false, nil
	}

	return s.client.HExists(context.Background(), *expiredTokens, token).Result()
}

func (s *InMemoryTokenStore) Remove(token string) error {
	tokenSruct, err := getToken(s.client, *existingTokens, token)
	if err != nil {
		return err
	}

	ctx := context.Background()
	s.client.HDel(ctx, *existingTokens, token).Err()
	if err != nil {
		return err
	}

	return s.client.HSet(ctx, *expiredTokens, token, tokenSruct).Err()
}

func (s *InMemoryTokenStore) Refesh(token string, lifetime time.Duration) (string, error) {
	tokenSruct, err := getToken(s.client, *existingTokens, token)
	if err != nil {
		return "", err
	}
	err = s.Remove(token)
	if err != nil {
		return "", err
	}
	return s.NewToken(tokenSruct.UserId, tokenSruct.Admin, lifetime)
}

func (s *InMemoryTokenStore) GetExistingToken(token string) (*models.AuthToken, error) {
	return getToken(s.client, *existingTokens, token)
}

func getToken(c *redis.Client, key string, token string) (*models.AuthToken, error) {
	authToken := &models.AuthToken{}
	bytes, err := c.HGet(context.Background(), *existingTokens, token).Bytes()
	if err != nil {
		return authToken, err
	}
	err = json.Unmarshal(bytes, authToken)
	return authToken, err
}
