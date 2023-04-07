package models

import (
	"encoding/json"
	"time"
)

type AuthToken struct {
	Token     string
	Admin     bool
	CreatedAt time.Time
	ExpiresAt time.Time
}

func (a *AuthToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

// func (a *AuthToken) BinaryUnmarshaler(data []byte) error {
// 	return json.Unmarshal(data, a)
// }

type TokenStore interface {
	NewToken(lifetime time.Duration) (string, error)
	IsExisting(token string) (bool, error)
	IsExpired(token string) (bool, error)
	Refesh(token string, lifetime time.Duration) (string, error)
	Remove(token string) error
}
