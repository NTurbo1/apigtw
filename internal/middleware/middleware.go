package middleware

import (
	"net/http"

	"github.com/nturbo1/reverse-proxy/internal/log"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Fixme("IMPLEMENT AUTHENTICATION/AUTHORIZATION!!!")

		next.ServeHTTP(w, r)
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Fixme("IMPLEMENT THE LOGGING MIDDLEWARE!!!")

		next.ServeHTTP(w, r)
	})
}

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Fixme("IMPLEMENT THE RATE LIMITER MIDDLEWARE!!!")

		next.ServeHTTP(w, r)
	})
}
