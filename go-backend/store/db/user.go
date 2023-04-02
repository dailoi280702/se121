package db_store

import (
	"github.com/dailoi280702/se121/go_backend/internal/utils"
	"github.com/dailoi280702/se121/go_backend/models"
)

type DbUserStore struct{}

func NewDbUserStore() *DbUserStore {
	return &DbUserStore{}
}

func (s *DbUserStore) GetUser(id string) (models.User, error) {
	return models.User{}, utils.UnplementedError
}

func (s *DbUserStore) AddUser(user models.User) error {
	return utils.UnplementedError
}

func (s *DbUserStore) UpdateUser(user models.User) error {
	return utils.UnplementedError
}

func (s *DbUserStore) DeleteUser(user models.User) error {
	return utils.UnplementedError
}

func (s *DbUserStore) GetUserByEmailOrName(name string, email string) (models.User, error) {
	return models.User{}, utils.UnplementedError
}
