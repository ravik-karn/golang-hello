package main

import (
	"time"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"personal/hello/handler"
)

func main() {
	n := negroni.New()
	router := handler.Router()
	n.UseHandler(router)

	s := &http.Server{
		Addr:           ":5400",
		Handler:        n,
		ReadTimeout:    200 * time.Second,
		WriteTimeout:   200 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
