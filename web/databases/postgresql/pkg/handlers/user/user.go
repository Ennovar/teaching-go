package user

import (
	"net/http"
	"time"
	"fmt"

	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/handlers"
	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/user"
)

func Login(w http.ResponseWriter, req *http.Request) {
	handlers.ExecuteTemplate(w, req, handlers.Data{
		PageTitle: "Login",
	}, "templates/content/user/login.html")
}

func Register(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()

		if ok := handlers.ParseFormValues([]string{"email", "password"}, req); !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		u, err := user.NewInstance(&user.User{
			Email:       req.Form["email"][0],
			Password:    req.Form["password"][0],
			Permissions: 0,
		})
		if err != nil {
			handlers.ExecuteTemplate(w, req, handlers.Data{
				PageTitle:     "Register",
				StatusMessage: fmt.Sprintf("An error has occurred: %v", err.Error()),
			}, "templates/content/user/register.html")
			return
		}

		err = u.Register()
		if err != nil {
			handlers.ExecuteTemplate(w, req, handlers.Data{
				PageTitle:     "Register",
				StatusMessage: fmt.Sprintf("An error has occurred: %v", err.Error()),
			}, "templates/content/user/register.html")
			return
		}

		req.URL.Path = "/user/login"
		handlers.ExecuteTemplate(w, req, handlers.Data{
			PageTitle:     "Login",
			StatusMessage: fmt.Sprintf("You have successfully registered under the email %v. You may now login.", u.Email),
		}, "templates/content/user/login.html")

		return
	}

	handlers.ExecuteTemplate(w, req, handlers.Data{
		PageTitle: "Register",
	}, "templates/content/user/register.html")
}

func Dashboard(w http.ResponseWriter, req *http.Request) {
	handlers.ExecuteTemplate(w, req, handlers.Data{
		PageTitle: "Dashboard",
	}, "templates/content/user/dashboard.html")
}

func Logout(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "postgres-user",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	handlers.ExecuteTemplate(w, req, handlers.Data{
		PageTitle:     "Logout",
		StatusMessage: "Successfully logged out.",
	}, "templates/content/user/logout.html")
}
