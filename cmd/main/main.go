package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcossabatini7/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":4000", r))
}
