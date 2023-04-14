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
	UserNameExistedErr              = errors.New("this username is already used")
	EmailExistedErr                 = errors.New("this email is already used")
	ErrIncorrectNameEmailOrPassword = errors.New("user name, email or password is not correct")
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
	ee := ErrExistedFields{}

	isNameExisted, err := ExistInDB(s.db, isUsernameExistedSql, user.Name)
	if err != nil {
		return err
	}
	if isNameExisted {
		ee.FieldNames = append(ee.FieldNames, "name")
	}

	if user.Email != "" {
		isEmailExisted, err := ExistInDB(s.db, isEmailExistedSql, user.Email)
		if err != nil {
			return err
		}
		if isEmailExisted {
			ee.FieldNames = append(ee.FieldNames, "email")
		}
	}
	log.Printf("\n%+v\n", user)
	log.Printf("\n%+v\n", ee)
	if len(ee.FieldNames) > 0 {
		return &ee
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

func (s *DbUserStore) VerifyUser(nameOrEmail string, password string) error {
	verified, err := ExistInDB(s.db, verifyUserQuery, nameOrEmail, password)
	if err != nil {
		return err
	}
	if !verified {
		return ErrIncorrectNameEmailOrPassword
	}
	return nil
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

const verifyUserQuery = `
    SELECT true FROM users WHERE (name = $1 OR email = $1) AND password = $2
    `
