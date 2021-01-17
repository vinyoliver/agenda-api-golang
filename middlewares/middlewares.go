package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// SetMiddlewareLogger displays a info message of the API
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

// SetMiddlewareJSON set the application Content-Type
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func Runsbefore(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("before...")
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func Runsafter(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("after...")
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
