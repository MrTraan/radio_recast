package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func createTrackHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println(vars)
}

func listTrackHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
