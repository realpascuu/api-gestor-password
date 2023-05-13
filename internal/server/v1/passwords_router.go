package v1

import (
	"encoding/json"
	auth "gestorpasswordapi/internal/middleware"
	"gestorpasswordapi/pkg/passwords"
	"gestorpasswordapi/pkg/response"
	"net/http"

	"github.com/go-chi/chi"
)

type PasswordsRouter struct {
	Repository passwords.Repository
}

func (pr *PasswordsRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.With(auth.Authorizator).Post("/", pr.CreateHandler)
	return r
}

func (pr *PasswordsRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := ctx.Value("id").(int)

	defer r.Body.Close()

	var p passwords.Passwords
	err := json.NewDecoder((r.Body)).Decode(&p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	p.UserID = uint(id)
	err = pr.Repository.Create(ctx, &p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, r, http.StatusCreated, p)
}