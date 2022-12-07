package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = 8080

type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	//set up app config
	var app application
	app.Domain = "example.com"

	// read from cli
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to db
	dbConn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	var dbRepoImpl dbrepo.PostgresDBRepo

	dbRepoImpl.DB = dbConn
	app.DB = &dbRepoImpl
	defer app.DB.Connection().Close()

	// start api
	log.Println("starting web server at port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
