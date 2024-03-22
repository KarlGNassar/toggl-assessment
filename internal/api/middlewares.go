package api

import (
	"errors"
	"net/http"
	"os"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")

		if apiKey != os.Getenv("API_KEY") {
			Unauthorized(w, ErrorInvalidApiKey, errors.New("invalid or missing api key"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
