package main

import (
	"Go/pkg/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)


	port := os.Getenv("PORT")
	if port == "" {
		port = "9010"
	}

	log.Printf("Starting server on port %s...\n", port)
    log.Fatal(http.ListenAndServe("localhost:"+port, r))
}