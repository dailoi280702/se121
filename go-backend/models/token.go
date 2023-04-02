package models

import "time"

type AuthToken struct {
	token     string
	admin     bool
	createdAt time.Time
	expiresAt time.Time
}

type TokenStore interface {
	NewToken() string
	IsExisting(token string) bool
	IsExpired(token string) bool
	Refesh(token string) string
	Remove(token string) string
}
