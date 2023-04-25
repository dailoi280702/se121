package service

import (
	"context"
	"encoding/json"
	"flag"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var (
	existingTokens = flag.String("existing tokens key", "existingTokens", "key for hashset of existing auth tokens")
	expiredTokens  = flag.String("expired tokens key", "expiredTokens", "key for hashset of expired auth tokens")
)

type AuthToken struct {
	Token     string    `json:"token"`
	UserId    string    `json:"userId"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func (a *AuthToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *AuthToken) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

func (s *Service) NewToken(userId string, isAdmin bool, lifetime time.Duration) (string, error) {
	key := uuid.NewString()
	token := &AuthToken{
		Token:     key,
		UserId:    userId,
		Admin:     isAdmin,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(lifetime),
	}
	err := s.rdb.HSet(context.Background(), *existingTokens, key, token).Err()
	return key, err
}

func (s *Service) IsExisting(token string) (bool, error) {
	return s.rdb.HExists(context.Background(), *existingTokens, token).Result()
}

func (s *Service) IsExpired(token string) (bool, error) {
	existing, err := s.IsExisting(token)
	if err != nil {
		return false, err
	}

	if existing {
		authToken, err := getToken(s.rdb, *existingTokens, token)
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

	return s.rdb.HExists(context.Background(), *expiredTokens, token).Result()
}

func (s *Service) Remove(token string) error {
	tokenSruct, err := getToken(s.rdb, *existingTokens, token)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = s.rdb.HDel(ctx, *existingTokens, token).Err()
	if err != nil {
		return err
	}

	return s.rdb.HSet(ctx, *expiredTokens, token, tokenSruct).Err()
}

func (s *Service) Refesh(token string, lifetime time.Duration) (string, error) {
	tokenSruct, err := getToken(s.rdb, *existingTokens, token)
	if err != nil {
		return "", err
	}
	err = s.Remove(token)
	if err != nil {
		return "", err
	}
	return s.NewToken(tokenSruct.UserId, tokenSruct.Admin, lifetime)
}

func (s *Service) GetExistingToken(token string) (*AuthToken, error) {
	return getToken(s.rdb, *existingTokens, token)
}

func getToken(c *redis.Client, key string, token string) (*AuthToken, error) {
	authToken := &AuthToken{}
	bytes, err := c.HGet(context.Background(), *existingTokens, token).Bytes()
	if err != nil {
		return authToken, err
	}
	err = json.Unmarshal(bytes, authToken)
	return authToken, err
}
