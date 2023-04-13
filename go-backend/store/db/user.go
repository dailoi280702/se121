package db_store

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/dailoi280702/se121/go_backend/internal/utils"
	"github.com/dailoi280702/se121/go_backend/models"
)

var (
	UserNameExistedErr = errors.New("this username is already used")
	EmailExistedErr    = errors.New("this email is already used")
)

type ErrExistedFields struct {
	FieldNames []string
}

func (e *ErrExistedFields) Error() string {
	return strings.Join(e.FieldNames, " ")
}

type DbUserStore struct {
	db *sql.DB
}

func NewDbUserStore(db *sql.DB) *DbUserStore {
	return &DbUserStore{db: db}
}

func (s *DbUserStore) GetUser(id string) (*models.User, error) {
	return nil, utils.UnplementedError
}

func (s *DbUserStore) AddUser(user models.User) error {
	existed := struct {
		username bool
		email    bool
	}{false, false}
	if err := s.db.QueryRow(isUsernameExistedSql, user.Name).Scan(&existed.username); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			existed.username = false
		default:
			log.Printf("\nerror checking username: %+v: %s\n", user, err)
			return err
		}
	}
	if err := s.db.QueryRow(isEmailExistedSql, user.Email).Scan(&existed.email); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			existed.username = false
		default:
			log.Printf("\nerror checking email: %+v: %s\n", user, err)
			return err
		}
	}
	log.Println(existed.username)

	err := ErrExistedFields{}
	if existed.username {
		err.FieldNames = append(err.FieldNames, "name")
	}
	if existed.email {
		err.FieldNames = append(err.FieldNames, "email")
	}
	if len(err.FieldNames) > 0 {
		return &err
	}

	if err := s.db.QueryRow(addUserSql, user.Name, user.Email, user.Password).Err(); err != nil {
		log.Printf("\nerror creating user: %+v: %s\n", user, err)
		return err
	}
	return nil
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

const addUserSql = `
    INSERT INTO users (name, email, password)
    VALUES ($1, $2, $3)
    `

const isUsernameExistedSql = `
    SELECT true FROM users WHERE name = $1
    `

const isEmailExistedSql = `
    SELECT true FROM users WHERE email = $1
    `
