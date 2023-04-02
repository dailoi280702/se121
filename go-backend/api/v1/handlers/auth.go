package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dailoi280702/se121/go_backend/models"
	"github.com/dailoi280702/se121/go_backend/store/db"
	memory_store "github.com/dailoi280702/se121/go_backend/store/memory"
	"github.com/go-chi/chi/v5"
)

const TokenLifetime = 24 * 5 * time.Hour

type AuthHandler struct {
	userStore  models.UserStore
	tokenStore models.TokenStore
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userStore:  db_store.NewDbUserStore(),
		tokenStore: memory_store.NewInMemoryTokenStore(),
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
	// :TODO get token from cookie
	dumpToken := "token"

	existed, err := h.tokenStore.IsExisting(dumpToken)
	if err != nil {
		mustSendError(err, w)
		return
	}
	if !existed {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// :TODO delete token from cookie
	err = h.tokenStore.Remove(dumpToken)
	if err != nil {
		mustSendError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("signed out"))
}

func (h AuthHandler) refresh(w http.ResponseWriter, r *http.Request) {
	// :TODO get token from cookie
	dumpToken := "token"

	existed, err := h.tokenStore.IsExisting(dumpToken)
	if err != nil {
		mustSendError(err, w)
		return
	}
	if !existed {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// :TODO send new token to cookie
	_, err = h.tokenStore.Refesh(dumpToken, TokenLifetime)
	if err != nil {
		mustSendError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("refeshed"))
}

func mustSendError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
		log.Panic(err)
	}
}
