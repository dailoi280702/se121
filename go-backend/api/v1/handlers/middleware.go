package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/dailoi280702/se121/go-backend/models"
)

func MustBeAuthenticated(next http.HandlerFunc, tokenStore models.TokenStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(*cookieAuthToken)
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				log.Println(err)
				http.Error(w, "cookie not found", http.StatusUnauthorized)
			default:
				log.Println(err)
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			return
		}

		log.Println(c.Value)
		// dumpToken := "token"
		existed, err := tokenStore.IsExisting(c.Value)
		if err != nil {
			MustSendError(err, w)
			return
		}

		if !existed {
			log.Println(c.Value)
			http.Error(w, "token not found", http.StatusUnauthorized)
			// w.WriteHeader(http.StatusUnauthorized)
			return
		}

		expired, err := tokenStore.IsExpired(c.Value)
		if err != nil {
			MustSendError(err, w)
			return
		}

		if expired {
			http.Error(w, "token expired", http.StatusUnauthorized)
			// w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
