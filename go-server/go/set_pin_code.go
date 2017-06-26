package main

import (
	"net/http"
)

type Set_pin_code struct {

}

func UpdateGsmPinCode(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

