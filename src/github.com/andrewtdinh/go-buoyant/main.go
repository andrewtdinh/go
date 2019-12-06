package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type request struct {
	Endpoint string `json:"Endpoint"`
}

type allRequests []request

var requests = allRequests{
	{
		Endpoint: "/",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/")
}

func printEndpoint(w http.ResponseWriter, r *http.Request) {
	var newRequest request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/books", printEndpoint).Methods("GET")
	router.HandleFunc("/books/{id: [0-9]+}", printEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", router))
}
