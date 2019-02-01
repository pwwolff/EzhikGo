package store

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"SearchUnstressed",
		"GET",
		"/API/Search/{unstressed}",
		controller.SearchUnstressed,
	},
	Route{
		"GetText",
		"GET",
		"/API/GetText/{TextName}",
		controller.GetText,
	},
	Route{
		"GetWordById",
		"GET",
		"/API/GetWord/{WordId}",
		controller.GetWordById,
	},
}

// NewRouter function configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
