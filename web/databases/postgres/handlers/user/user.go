package user

import "net/http"

const CookieName = "generic-forum-user"

func Verify(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie(CookieName)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Need to verify cookie when applicable to determine
	// whether or not I can write status 204.

	w.WriteHeader(http.StatusUnauthorized)
}