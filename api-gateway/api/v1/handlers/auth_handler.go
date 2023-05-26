package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/dailoi280702/se121/api-gateway/internal/service/auth"
	"github.com/dailoi280702/se121/api-gateway/models"
	cached_store "github.com/dailoi280702/se121/api-gateway/store/cache"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
)

const (
	TokenLifetime = 24 * 5 * time.Hour
)

var cookieAuthToken = flag.String("cookieAuthToken", "authToken", "name of auth token for cookie")

type AuthHandler struct {
	tokenStore  models.TokenStore
	authService auth.AuthServiceClient
}

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

func (h AuthHandler) signIn(w http.ResponseWriter, r *http.Request) {
	var req auth.SignInReq
	var res *auth.SignInRes

	convertJsonApiToGrpc(
		w, r,
		func() error {
			var err error
			res, err = h.authService.SignIn(context.Background(), &req)
			return err
		},
		convertWithJsonReqData(&req),
		convertWithPostFunc(func() {
			token := res.GetToken()
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

			user := res.GetUser()
			if err := json.NewEncoder(w).Encode(user); err != nil {
				MustSendError(err, w)
				return
			}
		}),
		convertWithCustomCodes(map[codes.Code]func(){
			codes.NotFound: func() {
				w.WriteHeader(http.StatusUnauthorized)
			},
		}))
}

func (h AuthHandler) signUp(w http.ResponseWriter, r *http.Request) {
	var req auth.SignUpReq
	convertJsonApiToGrpc(
		w, r,
		func() error {
			_, err := h.authService.SignUp(context.Background(), &req)
			return err
		},
		convertWithJsonReqData(&req),
		convertWithPostFunc(func() {
			log.Println(w.Write([]byte("ok")))
		}))
}

func (h AuthHandler) signOut(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(*cookieAuthToken)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	convertJsonApiToGrpc(
		w, r, func() error {
			_, err := h.authService.SignOut(context.Background(), &auth.SignOutReq{Token: c.Value})
			return err
		}, convertWithPostFunc(func() {
			c.Value = ""
			c.MaxAge = -1
			c.Path = "/"
			http.SetCookie(w, c)

			if _, err := w.Write([]byte("signed out")); err != nil {
				MustSendError(err, w)
			}
		}))
}

func (h AuthHandler) refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(*cookieAuthToken)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	var res *auth.RefreshRes
	convertJsonApiToGrpc(
		w, r, func() error {
			res, err = h.authService.Refresh(context.Background(), &auth.RefreshReq{Token: c.Value})
			return err
		}, convertWithPostFunc(func() {
			c.Value = res.GetToken()
			c.Path = "/"
			http.SetCookie(w, c)

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(res.GetUser()); err != nil {
				log.Panic(err)
			}
		}))
}
