package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/user"
)

// This new type, Data, will be used to shuttle data
// from the server into the templates.
type Data struct {
	PageTitle string
	StatusMessage string
	ExtraData interface{}
}

// This function will effectively reuse a bulk of the
// code that it takes to rewrite every time a template is
// executed, mostly due to the amount of files that are
// just reused by default in every template, like the layout,
// navigation, etc. This function automatically checks for
// whether or not public or private versions of templates
// need to be executed.
func ExecuteTemplate(w http.ResponseWriter, req *http.Request, data Data, contentTemplate string) {
	var templateFiles = []string{
		"templates/layout.html",
		"templates/navigation/navigation.public.html",
		contentTemplate,
	}

	c, err := req.Cookie("postgres-user")
	if err == nil {
		// Cookie is split by : with their ID on the left
		// of it and the sessionID on the right
		expanded := strings.Split(c.Value, ":")

		if len(expanded) == 2 {
			u, err := user.Get(expanded[0])
			if err == nil && u.SessionID == expanded[1] {
				templateFiles = []string{
					"templates/layout.html",
					"templates/navigation/navigation.private.html",
					contentTemplate,
				}
			}
		}
	}

	t := template.Must(template.ParseFiles(templateFiles...))
	t.ExecuteTemplate(w, "layout", data)
}


// Function ParseFormValues intakes a slice of strings that
// contains the index of all required form values as well as
// the request object and checks to see if all required indexes
// exist with data in them. If one value doesn't exist then this
// function returns false. If they all exist, the function returns
// true. req.ParseForm() must be called before utilizing this function.
func ParseFormValues(requiredValues []string, req *http.Request) bool {
	for _, required := range requiredValues {
		v := strings.TrimSpace(req.FormValue(required))

		if v == "" {
			return false
		}
	}

	return true
}