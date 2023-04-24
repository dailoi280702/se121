package models

import (
	"encoding/json"
	"time"
)

type AuthToken struct {
	Token     string    `json:"token"`
	UserId    string    `json:"userId"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func (a *AuthToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *AuthToken) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

type TokenStore interface {
	NewToken(userId string, isAdmin bool, lifetime time.Duration) (string, error)
	IsExisting(token string) (bool, error)
	IsExpired(token string) (bool, error)
	Refesh(token string, lifetime time.Duration) (string, error)
	Remove(token string) error
	GetExistingToken(token string) (*AuthToken, error)
}
