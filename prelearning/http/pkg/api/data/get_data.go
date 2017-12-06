package data

import "net/http"

func Get(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Secret data from the server!"))
}
