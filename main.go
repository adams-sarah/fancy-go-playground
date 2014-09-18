package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	setRoutes()

	http.ListenAndServe(":"+port, nil)
}

func setRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("HEAD", "GET").Name("HomeHandler")

	http.Handle("/", r)
}
