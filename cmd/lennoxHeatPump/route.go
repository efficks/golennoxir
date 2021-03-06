package main

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
		"State",
		"GET",
		"/state",
		GetState,
	},
	Route{
		"OffState",
		"PUT",
		"/off",
		PutOffState,
	},
	Route{
		"CoolState",
		"PUT",
		"/cool",
		PutCoolState,
	},
	Route{
		"DryState",
		"PUT",
		"/dry",
		PutDryState,
	},
	/*Route{
		"AutoState",
		"PUT",
		"/auto",
		PutAutoState,
	},*/
	Route{
		"HeatState",
		"PUT",
		"/heat",
		PutHeatState,
	},
	Route{
		"FanState",
		"PUT",
		"/fan",
		PutFanState,
	},
}
