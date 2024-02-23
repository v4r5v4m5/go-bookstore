package routes

import (
	"github.com/gorilla/mux"
	"github.com/v4r5v4m5/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")           // creates new book
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")               // get book
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")   // get book by id
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")    // update book
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE") // delete book
}
