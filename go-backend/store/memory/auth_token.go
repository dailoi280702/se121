package memory_store

import (
	"errors"
	"time"

	"github.com/dailoi280702/se121/go_backend/models"
	"github.com/google/uuid"
)

// An implementation of TokenStore which its data are stored in machine memory
type InMemoryTokenStore struct {
	existingTokens map[string]models.AuthToken
	expiredTokens  map[string]models.AuthToken
}

func NewInMemoryTokenStore() *InMemoryTokenStore {
	return &InMemoryTokenStore{
		existingTokens: map[string]models.AuthToken{},
		expiredTokens:  map[string]models.AuthToken{},
	}
}

func (s *InMemoryTokenStore) NewToken(lifetime time.Duration) (string, error) {
	newToken := uuid.NewString()
	s.existingTokens[newToken] = models.AuthToken{
		Token:     newToken,
		Admin:     false,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(lifetime),
	}

	return newToken, nil
}

func (s *InMemoryTokenStore) IsExisting(token string) (bool, error) {
	_, ok := s.existingTokens[token]
	return ok, nil
}

func (s *InMemoryTokenStore) IsExpired(token string) (bool, error) {
	_, ok := s.expiredTokens[token]
	return ok, nil
}

func (s *InMemoryTokenStore) Remove(token string) error {
	authToken, ok := s.existingTokens[token]
	if !ok {
		return errors.New("no token to be remove")
	}

	s.expiredTokens[token] = authToken
	delete(s.existingTokens, token)

	return nil
}

func (s *InMemoryTokenStore) Refesh(token string, lifetime time.Duration) (string, error) {
	err := s.Remove(token)
	if err != nil {
		return "", err
	}

	return s.NewToken(lifetime)
}
