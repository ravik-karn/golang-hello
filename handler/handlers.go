package handler

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Router() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", helloHandler()).Methods("GET")
	return r
}
