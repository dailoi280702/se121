package memory_store

import (
	"time"

	"github.com/dailoi280702/se121/go_backend/models"
)

// An implementation of TokenStore which its data are stored in machine memory
type InMemoryTokenStore struct {
	existingTokens []models.AuthToken
	existedTokens  []models.AuthToken
}

func (s *InMemoryTokenStore) NewToken(lifetime time.Time) (string, error) {
	return "", nil
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
