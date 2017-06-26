package main

import (
	"net/http"
	"fmt"
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
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/leila.ouederni/GridEye/1.0/",
		Index,
	},

	Route{
		"GetAlarmMonitoringEvent",
		"GET",
		"/leila.ouederni/GridEye/1.0/monitoring_events/alarm",
		GetAlarmMonitoringEvent,
	},

	Route{
		"GetGsmParam",
		"GET",
		"/leila.ouederni/GridEye/1.0/gsm_param",
		GetGsmParam,
	},

	Route{
		"GetNetworkParam",
		"GET",
		"/leila.ouederni/GridEye/1.0/network_param",
		GetNetworkParam,
	},

	Route{
		"GetMonMeasureMonitoringEvent",
		"GET",
		"/leila.ouederni/GridEye/1.0/monitoring_events/mon_measure",
		GetMonMeasureMonitoringEvent,
	},

	Route{
		"SetMonMsg",
		"PUT",
		"/leila.ouederni/GridEye/1.0/mon_msg",
		SetMonMsg,
	},

	Route{
		"GetMonitoringEvents",
		"GET",
		"/leila.ouederni/GridEye/1.0/monitoring_events",
		GetMonitoringEvents,
	},

	Route{
		"GetRegMeasureMonitoringEvent",
		"GET",
		"/leila.ouederni/GridEye/1.0/monitoring_events/reg_measure",
		GetRegMeasureMonitoringEvent,
	},

	Route{
		"UpdateGsmApn",
		"PUT",
		"/leila.ouederni/GridEye/1.0/gsm_param/apn",
		UpdateGsmApn,
	},

	Route{
		"UpdateNetwork",
		"PUT",
		"/leila.ouederni/GridEye/1.0/network_param",
		UpdateNetwork,
	},

	Route{
		"UpdateGsmPinCode",
		"PUT",
		"/leila.ouederni/GridEye/1.0/gsm_param/pin_code",
		UpdateGsmPinCode,
	},

}