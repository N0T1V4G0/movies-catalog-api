package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = 8080

type application struct {
	Domain string
}

func main() {
	//set up app config
	var app application
	app.Domain = "example.com"

	// read from cli

	// connect to db

	// start api
	log.Println("starting web server at port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
