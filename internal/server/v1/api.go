package v1

import (
	"gestorpasswordapi/internal/data"
	"net/http"

	"github.com/go-chi/chi"
)

func defineUserRoutes(r *chi.Mux) {
	ur := &UserRouter{
		Repository: &data.UserRepository{Data: data.New()},
	}

	r.Mount("/users", ur.Routes())
}

func definePasswordsRoutes(r *chi.Mux) {
	ur := &PasswordsRouter{
		Repository: &data.PasswordsRepository{Data: data.New()},
	}

	r.Mount("/passwords", ur.Routes())
}

func New() http.Handler {
	r := chi.NewRouter()

	defineUserRoutes(r)
	definePasswordsRoutes(r)
	return r
}
