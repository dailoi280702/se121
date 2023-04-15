package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func MustSendError(err error, w http.ResponseWriter) {
	log.Printf("\nBitch, you got an err: %s\n", err.Error())
	http.Error(w, fmt.Sprintf("server errror: %s", err.Error()), http.StatusInternalServerError)
}
