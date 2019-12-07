package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type request string

var allRequests []string

// var requests = allRequests

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/")
}

func printEndpoint(w http.ResponseWriter, r *http.Request) {
	newRequest := r.URL.Path
	fmt.Fprintf(w, "%v\n", newRequest)
	allRequests = append(allRequests, newRequest)
}

func getAllRequests(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(allRequests)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/requests", getAllRequests)
	router.HandleFunc("/books", printEndpoint).Methods("GET")
	router.HandleFunc("/books/{id}", printEndpoint).Methods("GET")
	router.HandleFunc("/books/new", printEndpoint).Methods("GET")
	router.HandleFunc("/books/{id}/edit", printEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", router))
}
