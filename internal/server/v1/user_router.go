package v1

import (
	"encoding/json"
	"fmt"
	auth "gestorpasswordapi/internal/middleware"
	"gestorpasswordapi/pkg/claim"
	"gestorpasswordapi/pkg/response"
	"gestorpasswordapi/pkg/user"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

type UserRouter struct {
	Repository user.Repository
}

func (ur *UserRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", ur.CreateHandler)
	r.With(auth.Authorizator).Put("/", ur.UpdateHandler)
	r.Post("/login", ur.LoginHandler)
	// ? r.Get("/{id}", ur.GetOneHandler)
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

	salt, err := u.GenerateRandomSalt()
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	u.Salt = salt

	password := u.GeneratePasswordHash(u.Password)
	u.Password = password

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

func (ur *UserRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := ctx.Value("id").(int)
	/* id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	} */
	defer r.Body.Close()

	var u user.User
	err := json.NewDecoder((r.Body)).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	salt, err := u.GenerateRandomSalt()
	if err != nil {
		response.HTTPError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	u.Salt = salt

	password := u.GeneratePasswordHash(u.Password)
	u.Password = password

	err = ur.Repository.Update(ctx, uint(id), u)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

func (ur *UserRouter) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder((r.Body)).Decode(&u)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	storedUser, err := ur.Repository.GetByEmail(ctx, u.Email)
	if err != nil {
		log.Fatal(err.Error())
		response.HTTPError(w, r, http.StatusNotFound, "Incorrect email or password")
		return
	}
	if !storedUser.PasswordMatch(u.Password) {
		response.HTTPError(w, r, http.StatusBadRequest, "Incorrect email or password")
		return
	}

	c := claim.Claim{ID: int(storedUser.ID), ExpDate: time.Now().Add(time.Hour * time.Duration(6)).Unix()}
	signingKey := os.Getenv("SIGNING_KEY")
	token, err := c.GetToken(signingKey)
	if err != nil {
		log.Fatal(err.Error())
		response.HTTPError(w, r, http.StatusInternalServerError, "Error while checking token")
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"token": token})
}
