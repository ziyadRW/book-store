package main

import (
	"github.com/gorilla/mux"
	"github.com/ziyadrw/go-book-store/internal/routes"
	_ "gorm.io/driver/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
