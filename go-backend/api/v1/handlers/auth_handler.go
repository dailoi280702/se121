package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/dailoi280702/se121/go_backend/internal/service/user"
	"github.com/dailoi280702/se121/go_backend/internal/utils"
	"github.com/dailoi280702/se121/go_backend/models"
	"github.com/dailoi280702/se121/go_backend/store/cache"
	"github.com/dailoi280702/se121/go_backend/store/db"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

const (
	TokenLifetime = 24 * 5 * time.Hour
	UsernameRegex = "^[A-Za-z0-9]+(?:[ _-][A-Za-z0-9]+)*$"
	EmailRegex    = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

var cookieAuthToken = flag.String("cookieAuthToken", "authToken", "name of auth token for cookie")

type AuthHandler struct {
	userStore   models.UserStore
	tokenStore  models.TokenStore
	userService user.UserServiceClient
}

// :TODO serialize password
func NewAuthHandler(redisClient *redis.Client, db *sql.DB, userService user.UserServiceClient) *AuthHandler {
	return &AuthHandler{
		userStore:   db_store.NewDbUserStore(db),
		tokenStore:  cached_store.NewRedisAuthTokenStore(redisClient),
		userService: userService,
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

// :TODO handle already authenticated request
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
			messages.Details.NameOrEmail = "neither user name is password are valid"
			valid = false
		}
	}
	if user.Password == "" {
		messages.Details.Password = "password cannot be empty"
		valid = false
	}

	// verify user
	var data *models.User
	if valid {
		data, err = h.userStore.VerifyUser(user.NameOrEmail, user.Password)
		if err != nil {
			switch {
			case errors.Is(err, db_store.ErrIncorrectNameEmailOrPassword):
				messages.Messages = append(messages.Messages, err.Error())
				valid = false
			default:
				MustSendError(err, w)
				return
			}
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
	token, err := h.tokenStore.NewToken(data.Id, data.IsAdmin, TokenLifetime)
	if err != nil {
		MustSendError(err, w)
		return
	}
	c := http.Cookie{
		Name:     *cookieAuthToken,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Expires:  time.Now().Add(TokenLifetime),
	}
	http.SetCookie(w, &c)

	// send user information
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Panic(err)
		return
	}
}

type signUpForm struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
}

func (h AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
	// get input
	user := signUpForm{}
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

	messages := struct {
		Messages []string          `json:"messages"`
		Details  map[string]string `json:"details"`
	}{
		Messages: []string{},
		Details:  map[string]string{},
	}

	// validate input
	if err := utils.ValidateField("name", user.Name, true, regexp.MustCompile(UsernameRegex)); err != nil {
		messages.Details["name"] = err.Error()
	}
	if err := utils.ValidateField("email", user.Email, false, regexp.MustCompile(EmailRegex)); err != nil {
		messages.Details["email"] = err.Error()
	}
	if err := utils.ValidateField("password", user.Password, true, nil); err != nil {
		messages.Details["password"] = err.Error()
	}
	if err := utils.ValidateField("", user.RePassword, true, nil); err != nil {
		messages.Details["rePassword"] = "please confirm password"
	} else if user.Password != user.RePassword {
		messages.Details["rePassword"] = "those password do not match"
	}

	w.Header().Set("Content-Type", "application/json")
	if len(messages.Details) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(messages); err != nil {
			log.Panic(err)
			return
		}
		return
	}

	// register user
	err = h.userStore.AddUser(models.User{Name: user.Name, Email: user.Email, Password: user.Password})
	if err != nil {
		var ee *db_store.ErrExistedFields

		switch {
		case errors.As(err, &ee):
			for _, field := range ee.FieldNames {
				messages.Details[field] = field + " is already used"
			}
			if len(messages.Details) != 0 {
				w.WriteHeader(http.StatusBadRequest)
				if err := json.NewEncoder(w).Encode(messages); err != nil {
					log.Panic(err)
					return
				}
				return
			}
		default:
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
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
	c.Path = "/"
	http.SetCookie(w, c)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("signed out"))
}

func (h AuthHandler) refresh(w http.ResponseWriter, r *http.Request) {
	// get token
	c, err := r.Cookie(*cookieAuthToken)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// refresh token
	token, err := h.tokenStore.Refesh(c.Value, TokenLifetime)
	if err != nil {
		MustSendError(err, w)
		return
	}

	c.Value = token
	c.Path = "/"
	http.SetCookie(w, c)

	// get and send user data
	// :TODO handle user that was deleted
	tokenData, err := h.tokenStore.GetExistingToken(token)
	if err != nil {
		MustSendError(err, w)
		return
	}

	// user, err := h.userStore.GetUser(tokenData.UserId)
	// if err != nil {
	// 	MustSendError(err, w)
	// 	return
	// }

	res, err := h.userService.GetUser(context.Background(), &user.GetUserReq{Id: tokenData.UserId})
	if err != nil {
		MustSendError(err, w)
		return
	}
	if res.User == nil {
		http.Error(w, "no user found", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res.User); err != nil {
		log.Panic(err)
		return
	}
}
