package index

import (
	"net/http"
	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/handlers"
)

func Index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	handlers.ExecutePublicTemplate(w, handlers.Data{
		PageTitle: "Home",
	}, "templates/content/index.html")
}
