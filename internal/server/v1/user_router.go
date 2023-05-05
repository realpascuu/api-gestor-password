package v1

import (
	"encoding/json"
	"fmt"
	"gestorpasswordapi/pkg/response"
	"gestorpasswordapi/pkg/user"
	"net/http"

	"github.com/go-chi/chi"
)

type UserRouter struct {
	Repository user.Repository
}

func (ur *UserRouter) Routes() http.Handler {
	r := chi.NewRouter()

	// TODO: add routes
	r.Post("/", ur.CreateHandler)
	// r.Get("/{id}", ur.GetOneHandler)
	return r
}

func (ur *UserRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder((r.Body)).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	_, noErr := ur.Repository.GetByEmail(ctx, u.Email)
	if noErr == nil {
		response.HTTPError(w, r, http.StatusBadRequest, "Email already exists")
		return
	}

	err = u.GenerateRandomSalt()
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	err = ur.Repository.Create(ctx, &u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	u.Password = ""
	u.Salt = ""
	w.Header().Add("Location", fmt.Sprintf("%s/%d", r.URL.String(), u.ID))
	response.JSON(w, r, http.StatusCreated, u)
}
