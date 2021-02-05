package agenda

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter returns an http handler capable of responding to http requests
func NewRouter(svc *Service) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(
		"/contact",
		loggingMiddleware(jsonMiddleware(svc.getContacts)),
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/contact",
		loggingMiddleware(jsonMiddleware(svc.addContact)),
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/contact/{id}",
		loggingMiddleware(jsonMiddleware(svc.getByID)),
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/contact/{id}",
		loggingMiddleware(jsonMiddleware(svc.updateContact)),
	).Methods(http.MethodPut)

	router.HandleFunc(
		"/contact/{id}",
		loggingMiddleware(jsonMiddleware(svc.deleteContact)),
	).Methods(http.MethodDelete)

	return router
}
