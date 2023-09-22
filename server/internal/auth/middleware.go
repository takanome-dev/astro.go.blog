package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/takanome-dev/blog-with-astro-golang/pkg/utils"
)

type MiddlewareFunc func(http.Handler) http.Handler

// Middleware takes Handler funcs and chains them to the main handler.
func Middleware(handler http.Handler, middlewares ...MiddlewareFunc) http.Handler {
  // The loop is reversed so the middlewares gets executed in the same
  // order as provided in the array.
  for i := len(middlewares); i > 0; i-- {
      handler = middlewares[i-1](handler)
  }
  return handler
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		defer log.Printf("[%s] %s %s: total time %s", r.Host, r.Method, r.URL, time.Since(start))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteError(w, errors.New("auth header is missing"), http.StatusUnauthorized)
			return
		}

		token := strings.Split(authHeader, " ")[1]
		if token == "" {
			utils.WriteError(w, errors.New("auth token is missing"), http.StatusUnauthorized)
			return
		}

		ok, err := utils.DecodeJwt(token)
		if !ok {
			utils.WriteError(w, err, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}