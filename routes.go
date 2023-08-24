package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func initRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "System was not implemented!")
	})

	return r
}
