package handlers

import (
	"html/template"
	"net/http"
)

// This new type, Data, will be used to shuttle data
// from the server into the templates.
type Data struct {
	PageTitle string
	ExtraData interface{}
}

// This function will effectively reuse a bulk of the
// code that it takes to rewrite every time a template is
// executed, mostly due to the amount of files that are
// just reused by default in every template, like the layout,
// navigation, etc.
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
