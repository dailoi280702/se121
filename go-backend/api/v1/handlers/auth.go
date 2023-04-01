package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h AuthHandler) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", h.refresh)
	router.Post("/", h.signIn)
	router.Put("/", h.signUp)
	router.Delete("/", h.signOut)

	return router
}

func (h AuthHandler) signIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("signIn"))
}

func (h AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("signUp"))
}

func (h AuthHandler) signOut(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("signOut"))
}

func (h AuthHandler) refresh(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("refesh"))
}
