package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = 8080

type application struct {
	DSN    string
	Domain string
}

func main() {
	//set up app config
	var app application
	app.Domain = "example.com"

	// read from cli
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to db

	// start api
	log.Println("starting web server at port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
