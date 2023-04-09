package handlers

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/dailoi280702/se121/go_backend/internal/utils"
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

type signInForm struct {
	NameOrEmail string `json:"nameOrEmail"`
	Password    string `json:"password"`
}

func (h AuthHandler) signIn(w http.ResponseWriter, r *http.Request) {
	// get input
	user := signInForm{}
	err := utils.DecodeJSONBody(w, r, &user)
	if err != nil {
		var mr *utils.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			MustSendError(err, w)
		}
		return
	}

	valid := true
	messages := struct {
		Messages []string   `json:"messages"`
		Details  signInForm `json:"details"`
	}{
		Messages: []string{},
		Details: signInForm{
			"",
			"",
		},
	}

	// validate input
	if user.NameOrEmail == "" {
		messages.Details.NameOrEmail = "user name or email cannot be empty"
		valid = false
	} else {
		emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		usernameRegex := regexp.MustCompile("^[A-Za-z0-9]+(?:[ _-][A-Za-z0-9]+)*$")
		isEmail := emailRegex.MatchString(user.NameOrEmail)
		isUsername := usernameRegex.MatchString(user.NameOrEmail)
		if !isEmail && !isUsername {
			messages.Details.NameOrEmail = "neither user name nor password are valid"
			valid = false
		}
	}
	if user.Password == "" {
		messages.Details.Password = "password cannot be empty"
		valid = false
	}

	// :TODO verify user
	if valid {
		if false {
			messages.Messages = append(messages.Messages, "name, email or password is not correct")
			valid = false
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(messages); err != nil {
			log.Panic(err)
			return
		}
		return
	}

	// generate auth token
	token, err := h.tokenStore.NewToken(TokenLifetime)
	if err != nil {
		MustSendError(err, w)
		return
	}
	c := http.Cookie{Name: *cookieAuthToken, Value: token}
	http.SetCookie(w, &c)

	// :TODO send user information
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Panic(err)
		return
	}
}

func (h AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
	// :TODO register user

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
