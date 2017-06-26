package main

import (
	"net/http"
)

type Mon_measure struct {

}

func GetMonMeasureMonitoringEvent(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

