package models

import "time"

type AuthToken struct {
	Token     string
	Admin     bool
	CreatedAt time.Time
	ExpiresAt time.Time
}

type TokenStore interface {
	NewToken(lifetime time.Duration) (string, error)
	IsExisting(token string) (bool, error)
	IsExpired(token string) (bool, error)
	Refesh(token string, lifetime time.Duration) (string, error)
	Remove(token string) error
}
