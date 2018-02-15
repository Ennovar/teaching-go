package about

import (
	"net/http"

	"github.com/Ennovar/teaching-go/web/templates_simple/pkg/handlers"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	handlers.ExecuteTemplate(w, handlers.Data{
		PageTitle: "About",
	}, "templates/about/index.html")
}

func Contact(w http.ResponseWriter, _ *http.Request) {
	handlers.ExecuteTemplate(w, handlers.Data{
		PageTitle: "Contact",
	}, "templates/about/contact.html")
}