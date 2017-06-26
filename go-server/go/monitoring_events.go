package main

import (
	"net/http"
)

type Monitoring_events struct {

}

func GetMonitoringEvents(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

