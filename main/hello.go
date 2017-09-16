package main

import (
	"github.com/codegangsta/negroni"
	"personal/hello/handler"
)

func main() {
	n := negroni.New()
	router := handler.Router()
	n.UseHandler(router)
	n.Run(":5400")
}
