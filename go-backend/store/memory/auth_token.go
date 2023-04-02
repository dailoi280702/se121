package memory_store

import (
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
		token:     newToken,
		admin:     false,
		createdAt: time.Now(),
		expiresAt: time.Now().Add(lifetime),
	}

	return newToken, nil
}

func (s *InMemoryTokenStore) IsExisting() (bool, error) {
	return false, nil
}

func (s *InMemoryTokenStore) IsExpired() (bool, error) {
	return false, nil
}

func (s *InMemoryTokenStore) Refesh(lifetime time.Time) (string, error) {
	return "", nil
}

func (s *InMemoryTokenStore) Remove() error {
	return nil
}
