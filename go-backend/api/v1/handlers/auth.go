package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dailoi280702/se121/go_backend/models"
	"github.com/dailoi280702/se121/go_backend/store/db"
	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	userStore models.UserStore
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userStore: db_store.NewDbUserStore(),
	}
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
	_, err := h.userStore.GetUser("id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Panic(err)
		}
		return
	}

	w.Write([]byte("signIn"))
}

func (h AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
	err := h.userStore.AddUser(models.User{})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Panic(err)
		}
		return
	}

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
