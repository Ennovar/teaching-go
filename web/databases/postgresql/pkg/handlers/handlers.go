package handlers

import (
	"net/http"
	"html/template"
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
// navigation, etc. This is specific to public templates, and
// will use public navigation bars, etc.
func ExecutePublicTemplate(w http.ResponseWriter, data Data, contentTemplate string) {
	templateFiles := []string {
		"templates/layout.html",
		"templates/navigation/navigation.public.html",
		contentTemplate,
	}

	t := template.Must(template.ParseFiles(templateFiles...))
	t.ExecuteTemplate(w, "layout", data)
}

// This function will effectively reuse a bulk of the
// code that it takes to rewrite every time a template is
// executed, mostly due to the amount of files that are
// just reused by default in every template, like the layout,
// navigation, etc. This is specific to private templates, and
// will use private navigation bars, etc.
func ExecutePrivateTemplate(w http.ResponseWriter, data Data, contentTemplate string) {
	templateFiles := []string {
		"templates/layout.html",
		"templates/navigation/navigation.private.html",
		contentTemplate,
	}

	t := template.Must(template.ParseFiles(templateFiles...))
	t.ExecuteTemplate(w, "layout", data)
}
