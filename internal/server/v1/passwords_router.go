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

	r.With(auth.Authorizator).Get("/{id}", pr.GetHandler)
	r.With(auth.Authorizator).Get("/", pr.GetAllHandler)
	r.With(auth.Authorizator).Post("/", pr.CreateHandler)
	r.With(auth.Authorizator).Delete("/{id}", pr.DeleteHandler)
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

func (pr *PasswordsRouter) GetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := ctx.Value("id").(int)

	defer r.Body.Close()

	idPassword := chi.URLParam(r, "id")
	p, err := pr.Repository.GetOne(ctx, idPassword)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, "Password not found")
		return
	}
	if p.UserID != uint(id) {
		response.HTTPError(w, r, http.StatusUnauthorized, "You are not authorized")
		return
	}
	response.JSON(w, r, http.StatusOK, p)
}

func (pr *PasswordsRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := ctx.Value("id").(int)

	defer r.Body.Close()

	p, err := pr.Repository.GetAll(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, "An error has ocurred")
		return
	}

	response.JSON(w, r, http.StatusOK, p)
}

func (pr *PasswordsRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := ctx.Value("id").(int)

	defer r.Body.Close()

	idPassword := chi.URLParam(r, "id")
	p, err := pr.Repository.GetOne(ctx, idPassword)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, "Password not found")
		return
	}
	if p.UserID != uint(id) {
		response.HTTPError(w, r, http.StatusUnauthorized, "You are not authorized")
		return
	}

	err = pr.Repository.Delete(ctx, idPassword)
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, "An error has ocurred")
		return
	}

	response.JSON(w, r, http.StatusNoContent, "")
}
