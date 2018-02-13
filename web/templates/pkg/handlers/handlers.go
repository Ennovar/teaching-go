package handlers

import (
	"html/template"
	"net/http"
)

type Data struct {
	PageTitle string
	ExtraData interface{}
}

func ExecuteTemplate(w http.ResponseWriter, data Data, contentTemplate string) {
	templateFiles := []string {
		"templates/layout.html",
		"templates/reuse/navigation.html",
		"templates/reuse/footer.html",
		"templates/reuse/jumbotron.html",
		contentTemplate,
	}

	t := template.Must(template.ParseFiles(templateFiles...))
	t.ExecuteTemplate(w, "layout", data)
}
