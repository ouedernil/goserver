package main

import (
	"net/http"
)

type Set_apn struct {

}

func UpdateGsmApn(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

