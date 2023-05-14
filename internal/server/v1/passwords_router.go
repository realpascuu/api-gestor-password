package v1

import (
	"context"
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

	r.With(auth.Authorizator).With(pr.PasswordAuthenticator).Get("/{id}", pr.GetHandler)
	r.With(auth.Authorizator).Get("/", pr.GetAllHandler)
	r.With(auth.Authorizator).Post("/", pr.CreateHandler)
	r.With(auth.Authorizator).With(pr.PasswordAuthenticator).Delete("/{id}", pr.DeleteHandler)
	r.With(auth.Authorizator).With(pr.PasswordAuthenticator).Put("/{id}", pr.UpdateHandler)
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
	p := ctx.Value("password").(passwords.Passwords)

	defer r.Body.Close()

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
	p := ctx.Value("password").(passwords.Passwords)

	err := pr.Repository.Delete(ctx, p.ID)
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, "An error has ocurred")
		return
	}

	response.JSON(w, r, http.StatusNoContent, nil)
}

func (pr *PasswordsRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	p := ctx.Value("password").(passwords.Passwords)

	var updatedPassword passwords.Passwords
	err := json.NewDecoder((r.Body)).Decode(&updatedPassword)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	err = pr.Repository.Update(ctx, p.ID, updatedPassword)
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, "An error has ocurred")
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

func (pr *PasswordsRouter) PasswordAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		ctx = context.WithValue(ctx, "password", p)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
