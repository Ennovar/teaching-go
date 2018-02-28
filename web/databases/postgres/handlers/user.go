package handlers

import "net/http"

func VerifyUser(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}