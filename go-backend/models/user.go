package models

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
	Email    string `json:"email,omitempty"`
}

type UserStore interface {
	GetUser(id string) (*User, error)
	GetUserByEmailOrName(name string, email string) (User, error)
	AddUser(user User) error
	UpdateUser(user User) error
	DeleteUser(user User) error
}
