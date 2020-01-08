package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shubham1010/mysql-nginx-docker/my-project/api"
	"github.com/shubham1010/mysql-nginx-docker/my-project/dbConnections"
)

func main() {
	database, err := dbConnections.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &api.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
