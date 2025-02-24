package routes

import (
	"github.com/gorilla/mux"
	handlers "github.com/ziyadrw/go-book-store/internal/handlers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", handlers.DeleteBook).Methods("DELETE")
}
