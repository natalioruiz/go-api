package album

import (
	"net/http"

	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/album",
		controller.Index,
	},
	Route{
		"AddAlbum",
		"Post",
		"/album",
		controller.AddAlbum,
	},
	Route{
		"GetAlbum",
		"Get",
		"/album/{id}",
		controller.GetAlbum,
	},
	Route{
		"UpdateAlbum",
		"Put",
		"/album/{id}",
		controller.UpdateAlbum,
	},
	Route{
		"DeleteAlbum",
		"Delete",
		"/album/{id}",
		controller.DeleteAlbum,
	},
}

// NewRouter configures a router for the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
