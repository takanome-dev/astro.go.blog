package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body []byte 
}

func (rec *ResponseRecorder) WriteHeader(statusCode int) {
	rec.StatusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (rec *ResponseRecorder) Write(body []byte) (int, error) {
	rec.Body = body
	return rec.ResponseWriter.Write(body)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &ResponseRecorder{
			ResponseWriter: w,
			StatusCode: http.StatusOK,
		}

		start := time.Now()
		next.ServeHTTP(rec, r)
		duration := time.Since(start)

		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		logger := log.Info()

		if rec.StatusCode != http.StatusOK {
			logger = log.Error().Bytes("body", rec.Body)
		}

		logger.Str("protocol", "http").
		Str("method", r.Method).
		Str("path", r.RequestURI).
		Int("status_code", rec.StatusCode).
		Str("status_text", http.StatusText(rec.StatusCode)).
		Dur("duration", duration).
		Msg("received an HTTP request")
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

		decoded, err := utils.DecodeJwt(token)
		if err != nil {
			utils.WriteError(w, err, http.StatusUnauthorized)
			return
		}

		// add UserID to context
		ctx := utils.CtxWithValue[utils.JwtUser](r.Context(), utils.JwtUser{UserID: decoded})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}