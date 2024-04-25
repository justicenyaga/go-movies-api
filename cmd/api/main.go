package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
	DSN    string
}

func main() {
	// set application config
	var app application

	app.Domain = "example.com"

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=letmein dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to the database

	// start a web server
	log.Println("Starting application on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
