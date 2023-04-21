package service

import (
	"errors"
	"time"
)

var (
	UserNameExistedErr              = errors.New("this username is already used")
	EmailExistedErr                 = errors.New("this email is already used")
	ErrIncorrectNameEmailOrPassword = errors.New("user name, email or password is not correct")
)

type User struct {
	Id       string    `json:"id,omitempty"`
	Name     string    `json:"name"`
	Email    string    `json:"email,omitempty"`
	ImageUrl string    `json:"imageUrl,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
	IsAdmin  bool      `json:"isAdmin,omitempty"`
	Password string    `json:"password,omitempty"`
}

func (s *Service) GetUser(id string) (*User, error) {
	user, err := GetObjectFromDB[User](s.DB, getUserByIdQuery, id)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func (s *Service) AddUser(user User) error {
	errs := ValidationErrors{errors: make(map[string]string)}

	isNameExisted, err := ExistInDB(s.DB, isUsernameExistedSql, user.Name)
	if err != nil {
		return err
	}
	if isNameExisted {
		errs.errors["name"] = "Username already in used"
	}

	if user.Email != "" {
		isEmailExisted, err := ExistInDB(s.DB, isEmailExistedSql, user.Email)
		if err != nil {
			return err
		}
		if isEmailExisted {
			errs.errors["email"] = "Username already in used"
		}
	}
	if len(errs.errors) > 0 {
		return &errs
	}

	if err := s.DB.QueryRow(addUserSql, user.Name, user.Email, user.Password).Err(); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateUser(user User) error {
	return UnplementedError
}

func (s *Service) DeleteUser(user User) error {
	return UnplementedError
}

func (s *Service) GetUserByEmailOrName(name string, email string) (User, error) {
	return User{}, UnplementedError
}

func (s *Service) VerifyUser(nameOrEmail string, password string) (*User, error) {
	user, err := GetObjectFromDB[User](s.DB, verifyUserQuery, nameOrEmail, password)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrIncorrectNameEmailOrPassword
	}
	user.Password = ""
	return user, nil
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
    SELECT * FROM users WHERE (name = $1 OR email = $1) AND password = $2
    `

const getUserByIdQuery = `
    SELECT * FROM users WHERE id = $1
    `
