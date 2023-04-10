package handlers

import (
	// "encoding/json"
	// "log"
	"net/http"
)

func MustSendError(err error, w http.ResponseWriter) {
	// w.WriteHeader(http.StatusInternalServerError)
	//
	http.Error(w, err.Error(), http.StatusInternalServerError)
	// if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
	// 	log.Panic(err)
	// }
}
