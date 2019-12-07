package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type request struct {
	Endpoint string `json:"Endpoint"`
}

type allRequests []request

var requests = allRequests{}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/")
}

func printEndpoint(w http.ResponseWriter, r *http.Request) {
	var newRequest request
	newRequest["Endpoint"] := r.URL.Path
	fmt.Fprintf(w, "%+v\n", newRequest)
	requests = append(requests, newRequest)
}

func getAllRequests(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(requests)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/requests", getAllRequests)
	router.HandleFunc("/books", printEndpoint).Methods("GET")
	router.HandleFunc("/books/{id: [0-9]+}", printEndpoint).Methods("GET")
	router.HandleFunc("/books/new", printEndpoint).Methods("GET")
	router.HandleFunc("/books/{id: [0-9]+}/edit", printEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", router))
}
