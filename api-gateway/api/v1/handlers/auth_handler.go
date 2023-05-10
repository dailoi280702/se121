package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/dailoi280702/se121/api-gateway/internal/service/auth"
	"github.com/dailoi280702/se121/api-gateway/internal/utils"
	"github.com/dailoi280702/se121/api-gateway/models"
	cached_store "github.com/dailoi280702/se121/api-gateway/store/cache"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	TokenLifetime = 24 * 5 * time.Hour
)

var cookieAuthToken = flag.String("cookieAuthToken", "authToken", "name of auth token for cookie")

type AuthHandler struct {
	tokenStore  models.TokenStore
	authService auth.AuthServiceClient
}

// :TODO serialize password
func NewAuthHandler(redisClient *redis.Client, db *sql.DB, authService auth.AuthServiceClient) *AuthHandler {
	return &AuthHandler{
		tokenStore:  cached_store.NewRedisAuthTokenStore(redisClient),
		authService: authService,
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
	w.Header().Set("Content-Type", "application/json")
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

	req, err := h.authService.SignIn(context.Background(), &auth.SignInReq{
		NameOrEmail: input.NameOrEmail,
		Password:    input.Password,
	})
	if err != nil {
		code := status.Code(err)
		s, _ := status.FromError(err)
		switch code {
		case codes.InvalidArgument:
			w.WriteHeader(http.StatusBadRequest)
		case codes.NotFound:
			w.WriteHeader(http.StatusUnauthorized)
		case codes.Unavailable:
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		default:
			MustSendError(err, w)
			return
		}
		_, err = w.Write([]byte(s.Message()))
		if err != nil {
			MustSendError(err, w)
		}
		return
	}

	token := req.GetToken()
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

	user := req.GetUser()
	if err := json.NewEncoder(w).Encode(user); err != nil {
		MustSendError(err, w)
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

	// messages := struct {
	// 	Messages []string          `json:"messages"`
	// 	Details  map[string]string `json:"details"`
	// }{
	// 	Messages: []string{},
	// 	Details:  map[string]string{},
	// }

	// validate input
	// if err := utils.ValidateField("name", input.Name, true, regexp.MustCompile(UsernameRegex)); err != nil {
	// 	messages.Details["name"] = err.Error()
	// }
	// if err := utils.ValidateField("email", input.Email, false, regexp.MustCompile(EmailRegex)); err != nil {
	// 	messages.Details["email"] = err.Error()
	// }
	// if err := utils.ValidateField("password", input.Password, true, nil); err != nil {
	// 	messages.Details["password"] = err.Error()
	// }
	// if err := utils.ValidateField("", input.RePassword, true, nil); err != nil {
	// 	messages.Details["rePassword"] = "please confirm password"
	// } else if input.Password != input.RePassword {
	// 	messages.Details["rePassword"] = "those password do not match"
	// }
	//
	// w.Header().Set("Content-Type", "application/json")
	// if len(messages.Details) != 0 {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	if err := json.NewEncoder(w).Encode(messages); err != nil {
	// 		log.Panic(err)
	// 		return
	// 	}
	// 	return
	// }

	// register user
	_, err = h.authService.SignUp(context.Background(), &auth.SignUpReq{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		RePasssword: input.RePassword,
	})

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		code := status.Code(err)
		s, _ := status.FromError(err)

		switch code {
		case codes.InvalidArgument:
			w.WriteHeader(http.StatusBadRequest)
		case codes.AlreadyExists:
			w.WriteHeader(http.StatusConflict)
		case codes.Unavailable:
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		default:
			MustSendError(err, w)
			return
		}

		_, err = w.Write([]byte(s.Message()))
		if err != nil {
			MustSendError(err, w)
		}
		return
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

	_, err = h.authService.SignOut(context.Background(), &auth.SignOutReq{Token: c.Value})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.Unavailable:
			http.Error(w, "auth service unavailable", http.StatusServiceUnavailable)
		default:
			MustSendError(err, w)
		}
		return
	}

	c.Value = ""
	c.MaxAge = -1
	c.Path = "/"
	http.SetCookie(w, c)

	if _, err := w.Write([]byte("signed out")); err != nil {
		MustSendError(err, w)
	}
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
	// token, err := h.tokenStore.Refesh(c.Value, TokenLifetime)
	// if err != nil {
	// 	MustSendError(err, w)
	// 	return
	// }

	// get and send user data
	// :TODO handle user that was deleted
	// tokenData, err := h.tokenStore.GetExistingToken(token)
	// if err != nil {
	// 	code := status.Code(err)
	// 	switch code {
	// 	case codes.Unavailable:
	// 		http.Error(w, "auth service unavailable", http.StatusServiceUnavailable)
	// 	default:
	// 		MustSendError(err, w)
	// 	}
	// 	return
	// }
	//
	// res, err := h.userService.GetUser(context.Background(), &user.GetUserReq{Id: tokenData.UserId})
	// if err != nil {
	// 	code := status.Code(err)
	// 	switch code {
	// 	case codes.NotFound:
	// 		http.Error(w, "no user found", http.StatusNoContent)
	// 	case codes.Internal:
	// 		http.Error(w, "service unabailable", http.StatusServiceUnavailable)
	// 	default:
	// 		MustSendError(err, w)
	// 	}
	// 	return
	// }

	res, err := h.authService.Refresh(context.Background(), &auth.RefreshReq{Token: c.Value})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			http.Error(w, "no user found", http.StatusNoContent)
		case codes.Unavailable:
			http.Error(w, "service unabailable", http.StatusServiceUnavailable)
		default:
			MustSendError(err, w)
		}
		return
	}

	c.Value = res.GetToken()
	c.Path = "/"
	http.SetCookie(w, c)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res.GetUser()); err != nil {
		log.Panic(err)
	}
}
