package routes

import (
	"agenda-api/middlewares"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct
type Route struct {
	URI     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

// Load the routes
func Load() []Route {
	routes := contactRoutes
	//Add more roues
	// routes = append(routes, postsRoutes...)
	return routes
}

func HandleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range Load() {
		router.HandleFunc(route.URI,
			middlewares.SetMiddlewareLogger(
				middlewares.SetMiddlewareJSON(route.Handler),
			),
		).Methods(route.Method)

	}
	// return r
	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
