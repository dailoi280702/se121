package handlers

import (
	"net/http"

	"github.com/dailoi280702/se121/go_backend/models"
)

func MustBeAuthenticated(next http.HandlerFunc, tokenStore models.TokenStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dumpToken := "token"
		existed, err := tokenStore.IsExisting(dumpToken)
		if err != nil {
			MustSendError(err, w)
			return
		}

		if !existed {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		expired, err := tokenStore.IsExpired(dumpToken)
		if err != nil {
			MustSendError(err, w)
			return
		}

		if expired {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
