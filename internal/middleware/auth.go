package auth

import (
	"context"
	"errors"
	"gestorpasswordapi/pkg/claim"
	"gestorpasswordapi/pkg/response"
	"net/http"
	"os"
	"strings"
)

const startedHeader = "Bearer "

func Authorizator(next http.Handler) http.Handler {
	signingString := os.Getenv("SIGNING_KEY")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		tokenString, err := tokenFromAuthorization(authorization)
		if err != nil {
			response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
			return
		}

		c, err := claim.GetFromToken(tokenString, signingString)
		if err != nil {
			response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
			return
		}
		// * save info in the request
		// TODO: get user from context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", c.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("authorization is required")
	}

	if !strings.HasPrefix(authorization, startedHeader) {
		return "", errors.New("invalid authorization format")
	}

	authString := strings.Split(authorization, " ")
	if len(authString) != 2 {
		return "", errors.New("invalid authorization format")
	}

	return authString[1], nil
}
