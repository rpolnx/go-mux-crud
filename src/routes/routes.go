package routes

import (
	"main/src/controllers"

	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")

	return router;
}