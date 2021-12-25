package middleware

import (
	"context"
	"net/http"
)

var UserCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type User struct {
	Name string
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Token"))
			return
		}

		ctx := context.WithValue(
			context.Background(),
			UserCtxKey,
			&User{
				Name: "Gustavo Delfim",
			},
		)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
