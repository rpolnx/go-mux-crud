package main

import (
	"log"
	"net/http"

	"go.uber.org/zap"

	"main/src/routes"

	"github.com/gorilla/mux"
)

var logger, _ = zap.NewProduction()


func main() {
	router := mux.NewRouter()

	booksRouter:= routes.RegisterBookStoreRoutes()

	router.NewRoute().Handler(booksRouter)
	
	log.Fatal(http.ListenAndServe(":8080", router))
	
}