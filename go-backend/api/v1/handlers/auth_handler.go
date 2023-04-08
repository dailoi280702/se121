package handlers

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/dailoi280702/se121/go_backend/models"
	"github.com/dailoi280702/se121/go_backend/store/cache"
	"github.com/dailoi280702/se121/go_backend/store/db"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

const TokenLifetime = 24 * 5 * time.Hour

var cookieAuthToken = flag.String("cookieAuthToken", "authToken", "name of auth token for cookie")

type AuthHandler struct {
	userStore  models.UserStore
	tokenStore models.TokenStore
}

func NewAuthHandler(redisClient *redis.Client) *AuthHandler {
	return &AuthHandler{
		userStore:  db_store.NewDbUserStore(),
		tokenStore: cached_store.NewRedisAuthTokenStore(redisClient),
	}
}

func (h AuthHandler) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", MustBeAuthenticated(h.refresh, h.tokenStore))
	router.Post("/", h.signIn)
	router.Put("/", h.signUp)
	router.Delete("/", MustBeAuthenticated(h.signOut, h.tokenStore))

	return router
}

func (h AuthHandler) signIn(w http.ResponseWriter, r *http.Request) {
	// :TODO delete next line
	token, err := h.tokenStore.NewToken(TokenLifetime)
	if err != nil {
		MustSendError(err, w)
	}
	c := http.Cookie{Name: *cookieAuthToken, Value: token}
	http.SetCookie(w, &c)

	if err := json.NewEncoder(w).Encode(token); err != nil {
		log.Panic(err)
		return
	}
	return

	_, err = h.userStore.GetUser("id")
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
	c, err := r.Cookie(*cookieAuthToken)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	err = h.tokenStore.Remove(c.Value)
	if err != nil {
		MustSendError(err, w)
		return
	}

	c.Value = ""
	c.MaxAge = -1
	http.SetCookie(w, c)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("signed out"))
}

func (h AuthHandler) refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(*cookieAuthToken)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	token, err := h.tokenStore.Refesh(c.Value, TokenLifetime)
	if err != nil {
		MustSendError(err, w)
		return
	}

	c.Value = token
	http.SetCookie(w, c)

	// :TODO send user infomation
	if err = json.NewEncoder(w).Encode(token); err != nil {
		log.Panic(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("refeshed"))
}
