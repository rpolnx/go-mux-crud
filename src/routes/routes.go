package routes

import (
	"main/src/controllers"

	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes() *mux.Router {

	subRouter := mux.NewRouter().PathPrefix("/books").Subrouter()

	subRouter.HandleFunc("", controllers.CreateBook).Methods("POST")
	subRouter.HandleFunc("", controllers.GetBook).Methods("GET")
	subRouter.HandleFunc("/{bookId}", controllers.GetBookById).Methods("GET")
	subRouter.HandleFunc("/{bookId}", controllers.UpdateBook).Methods("PUT")
	subRouter.HandleFunc("/{bookId}", controllers.DeleteBook).Methods("DELETE")

	return subRouter
}