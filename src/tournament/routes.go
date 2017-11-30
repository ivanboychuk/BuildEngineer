package tournament

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"fund",
		"GET",
		"/fund",
		fundPlayer,
	},
	Route{
		"take",
		"GET",
		"/take",
		takePlayer,
	},
	Route{
		"joinTournament",
		"GET",
		"/joinTournament",
		joinTournament,
	},
	Route{
		"resultTournament",
		"POST",
		"/resultTournament",
		resultTournament,
	},
	Route{
		"announceTournament",
		"GET",
		"/announceTournament",
		announceTournament,
	},
	Route{
		"balance",
		"GET",
		"/balance",
		balance,
	},
}