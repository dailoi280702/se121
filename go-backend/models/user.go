package models

import (
	// "database/sql"
	"time"
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

type UserStore interface {
	GetUser(id string) (*User, error)
	GetUserByEmailOrName(name string, email string) (User, error)
	AddUser(user User) error
	UpdateUser(user User) error
	DeleteUser(user User) error
	VerifyUser(nameOrEmail string, password string) (*User, error)
}
