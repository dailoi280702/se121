package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/dailoi280702/se121/go_backend/internal/service/user"
	"github.com/dailoi280702/se121/go_backend/internal/utils"
	"github.com/dailoi280702/se121/go_backend/models"
	cached_store "github.com/dailoi280702/se121/go_backend/store/cache"
	db_store "github.com/dailoi280702/se121/go_backend/store/db"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	input := signInForm{}
	err := utils.DecodeJSONBody(w, r, &input)
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
	if input.NameOrEmail == "" {
		messages.Details.NameOrEmail = "user name or email cannot be empty"
		valid = false
	} else {
		emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		usernameRegex := regexp.MustCompile("^[A-Za-z0-9]+(?:[ _-][A-Za-z0-9]+)*$")
		isEmail := emailRegex.MatchString(input.NameOrEmail)
		isUsername := usernameRegex.MatchString(input.NameOrEmail)
		if !isEmail && !isUsername {
			messages.Details.NameOrEmail = "neither user name is password are valid"
			valid = false
		}
	}
	if input.Password == "" {
		messages.Details.Password = "password cannot be empty"
		valid = false
	}

	// verify user
	var data *user.User
	if valid {
		data, err = h.userService.VerifyUser(context.Background(), &user.VerifyUserReq{
			NameOrEmail: input.NameOrEmail,
			Passord:     input.Password,
		})
		if err != nil {
			code := status.Code(err)
			s, _ := status.FromError(err)
			switch code {
			case codes.NotFound:
				messages.Messages = append(messages.Messages, s.Message())
				valid = false
			case codes.Internal:
				http.Error(w, "service unabailable", http.StatusServiceUnavailable)
				return
			default:
				MustSendError(err, w)
				return
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
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
	input := signUpForm{}
	err := utils.DecodeJSONBody(w, r, &input)
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
	if err := utils.ValidateField("name", input.Name, true, regexp.MustCompile(UsernameRegex)); err != nil {
		messages.Details["name"] = err.Error()
	}
	if err := utils.ValidateField("email", input.Email, false, regexp.MustCompile(EmailRegex)); err != nil {
		messages.Details["email"] = err.Error()
	}
	if err := utils.ValidateField("password", input.Password, true, nil); err != nil {
		messages.Details["password"] = err.Error()
	}
	if err := utils.ValidateField("", input.RePassword, true, nil); err != nil {
		messages.Details["rePassword"] = "please confirm password"
	} else if input.Password != input.RePassword {
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
	stream, err := h.userService.CreateUser(context.Background(), &user.CreateUserReq{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Email,
	})
	if err != nil {
		MustSendError(err, w)
	}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			code := status.Code(err)
			switch code {
			case codes.InvalidArgument:
				w.WriteHeader(http.StatusBadRequest)
			case codes.AlreadyExists:
				w.WriteHeader(http.StatusConflict)
			case codes.Internal:
				http.Error(w, "service unabailable", http.StatusServiceUnavailable)
				return
			default:
				MustSendError(err, w)
				return
			}
			if err := json.NewEncoder(w).Encode(messages); err != nil {
				log.Panic(err)
			}
			return
		}

		messages.Details = req.GetErrors()
	}

	log.Println(w.Write([]byte("ok")))
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

	log.Panic(w.Write([]byte("signed out")))
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

	// get and send user data
	// :TODO handle user that was deleted
	tokenData, err := h.tokenStore.GetExistingToken(token)
	if err != nil {
		MustSendError(err, w)
		return
	}

	res, err := h.userService.GetUser(context.Background(), &user.GetUserReq{Id: tokenData.UserId})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			http.Error(w, "no user found", http.StatusNoContent)
		case codes.Internal:
			http.Error(w, "service unabailable", http.StatusServiceUnavailable)
		default:
			MustSendError(err, w)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res.User); err != nil {
		log.Panic(err)
	}

	c.Value = token
	c.Path = "/"
	http.SetCookie(w, c)
}
