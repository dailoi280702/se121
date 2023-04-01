package models

type User struct {
	id       string
	name     string
	password string
	imageUrl string
	email    string
}

type UserStore interface {
	GetUser(id string) (User, error)
	GetUserByEmailOrName(name string, email string) (User, error)
	AddUser(user User) error
	UpdateUser(user User) error
	DeleteUser(user User) error
}
