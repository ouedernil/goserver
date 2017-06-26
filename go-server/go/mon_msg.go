package main

import (
	"net/http"
)

type Mon_msg struct {

}

func SetMonMsg(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

