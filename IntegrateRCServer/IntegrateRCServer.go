package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Lauenburg/GoThings/IntegrateRCServer/oTemp"
	"github.com/gorilla/mux"
)

/*
IntegrateRCServer is executed with following command: ./IntegrateRCServer <devId> <appId> <appAccessKey>
The command line arguments are needed for the request: GET/temperature/"+appID+"/"+devID

The package "oTemper" provides the microservice "OTemp" for receiving the office temperature
You can find the package "oTemper" in "vendor/github.com/lauenburg/IntegrateRCServer/oTemp" within the "IntegrateRCServer" folder

In case that the browser request for the website icon ("favicon.ico") is to be counted in addition to the user's
request move the "counter++" outside of the if-else statement and omit the second else-if case.
*/

//Counts the Requests to the server
var counter int = 1

//Holds the device Id
var devID string
var appID string
var appAccessKey string

// function executed for request "/Count"
func returnCounter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are talking to instance %s:[internal_port]. This is the %dth request to this instance.", r.Host, counter)
	counter++
}

// function executed for request "/Count"
func returnTemperature(w http.ResponseWriter, r *http.Request) {
	uplinkMessage := oTemp.OTemp(devID, appID, appAccessKey)
	fmt.Fprintf(w, "It's currently %.2f degrees in office %s/%s", uplinkMessage.PayloadFields["temperature"], appID, devID)
	counter++
}

// function executed when the router can not find any matches
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing to see here. But the request still counts!")
	counter++
}

func main() {

	//Ensures that the user entered the necessary three arguments at execution
	if len(os.Args) != 4 {
		fmt.Println("Please make sure to provide the arguments devId, appId and appAccessKey when executing the programm.")
		os.Exit(1)
	}

	devID = os.Args[1]
	appID = os.Args[2]
	appAccessKey = os.Args[3]

	// Init router
	r := mux.NewRouter()

	// Route handle: Returning the counter and incrementing the counter
	r.HandleFunc("/Count", returnCounter).Methods("GET")

	// Route handle: Returning the temperature and incrementing the counter
	r.HandleFunc("/Temperature/"+appID+"/"+devID, returnTemperature).Methods("GET")

	//NotFoundHandler to catch error 404 page not found and increment counter
	r.NotFoundHandler = http.HandlerFunc(notFound)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
