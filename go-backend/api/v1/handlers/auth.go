package handlers

import "github.com/go-chi/chi/v5"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h AuthHandler) Routes() chi.Router {
	router := chi.NewRouter()

	return router
}
