package application

import (
	"io"
	"net/http"
	"os"
)

func (app *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	if len(path) == 0 {
		path = (app.Directory + app.DocumentRoot + "index.html")
	} else {
		path = (app.Directory + app.DocumentRoot + path)
	}

	api := app.handleAPI(res, req)
	if api {
		return // HTTP logic will be handled by API
	}

	f, err := os.Open(path)

	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	ct, err := app.getContentType(req)

	if err != nil {
		http.Error(res, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	res.Header().Add("Content-Type", ct)
	_, err = io.Copy(res, f)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
