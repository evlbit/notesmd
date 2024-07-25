package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/evlbit/notesmd/internal/auth"
	"github.com/evlbit/notesmd/internal/handlers"
)

type contextKey string

const UserKey contextKey = "userId"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getTokenFromRequest(r)
		if err != nil {
			permissionDenied(w)
			return
		}

		userId, err := auth.ValidateJWT(token)
		if err != nil {
			permissionDenied(w)
			return
		}

		// TODO: check if userId exists in database

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func getTokenFromRequest(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", fmt.Errorf("authorization header was not set")
	}

	tokenSplit := strings.Split(token, " ")
	if len(tokenSplit) != 2 {
		return "", fmt.Errorf("authorization header was not formed correctly")
	}

	if strings.TrimSpace(tokenSplit[0]) != "Bearer" {
		return "", fmt.Errorf("expected a Bearer token")
	}

	return strings.TrimSpace(tokenSplit[1]), nil
}

func permissionDenied(w http.ResponseWriter) {
	handlers.WriteError(w, 400, fmt.Errorf("permission denied"))
}
