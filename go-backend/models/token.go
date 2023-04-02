package models

import "time"

type AuthToken struct {
	token     string
	admin     bool
	createdAt time.Time
	expiresAt time.Time
}

type TokenStore interface {
	NewToken(lifetime time.Time) (string, error)
	IsExisting(token string) (bool, error)
	IsExpired(token string) (bool, error)
	Refesh(token string, lifetime time.Time) (string, error)
	Remove(token string) error
}
