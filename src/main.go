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
	mainRouter := mux.NewRouter()

	booksRouter := routes.RegisterBookStoreRoutes()

	mainRouter.Handle("/books", booksRouter)
	
	log.Fatal(http.ListenAndServe(":8080", mainRouter))
	
}