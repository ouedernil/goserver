package main

import (
	"net/http"
)

type Set_network_param struct {

}

func UpdateNetwork(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

