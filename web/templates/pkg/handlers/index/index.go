package index

import (
	"net/http"

	"github.com/Ennovar/teaching-go/web/templates/pkg/handlers"
)

func Index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)

		handlers.ExecuteTemplate(w, handlers.Data{
			PageTitle: "404",
		}, "templates/status/404.html")

		return
	}

	handlers.ExecuteTemplate(w, handlers.Data{
		PageTitle: "Home",
	}, "templates/index.html")
}