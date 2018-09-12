package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
The task states that the number of requests is to be counted and printed out on request (/Count).
*/

//Counts the Requests to the server
var counter int = 1

// function executed for request "/Count"
func returnCounter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are talking to instance %s:[internal_port]. This is the %dth request to this instance.", r.Host, counter)
	counter++
}

// function executed when the router can not find any matches
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing to see here. But the request still counts!")
	counter++
}

func main() {

	// Init router
	r := mux.NewRouter()

	// Route handle: Returning the counter and incrementing the counter
	r.HandleFunc("/count", returnCounter).Methods("GET")

	//NotFoundHandler to catch error 404 page not found and increment counter
	r.NotFoundHandler = http.HandlerFunc(notFound)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
