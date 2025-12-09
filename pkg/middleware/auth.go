package middleware

import (
	"context"
	"http_5/configs"
	"http_5/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	CtxEmailKey key = "CtxEmailKey"
)

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer") {
			writeUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			writeUnauthed(w)
			return
		}
		// Save email in context
		ctx := context.WithValue(r.Context(), CtxEmailKey, data.Email)
		// Put new context instead of original
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
