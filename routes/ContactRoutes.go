package routes

import (
	"agenda-api/controllers"
	"net/http"
)

var contactRoutes = []Route{
	Route{
		URI:     "/contact",
		Method:  http.MethodGet,
		Handler: controllers.GetContacts,
	},
	Route{
		URI:     "/contact",
		Method:  http.MethodPost,
		Handler: controllers.AddContact,
	},
	Route{
		URI:     "/contact/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetById,
	},
	Route{
		URI:     "/contact/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateContact,
	},
	Route{
		URI:     "/contact/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteContact,
	},
}
