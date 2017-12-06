package application

import (
	"errors"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

func (app *Application) getContentType(req *http.Request) (string, error) {
	var ct string

	path := req.URL.Path[1:]
	if len(path) == 0 {
		path = (app.Directory + app.DocumentRoot + "index.html")
	} else {
		path = (app.Directory + app.DocumentRoot + path)
	}

	path = filepath.Base(path)
	ext := filepath.Ext(path)
	fn := strings.TrimSuffix(path, ext)

	if ext == "" || fn == "" {
		return "", errors.New("Invalid file path")
	}

	ct = mime.TypeByExtension(ext)

	if ct == "" {
		return "", errors.New("Could not get content type")
	}

	return ct, nil
}
