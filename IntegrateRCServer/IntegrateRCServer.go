package main

import (
"fmt"
"log"
"net/http"
"os"
"github.com/lauenburg/IntegrateRCServer/oTemp"
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



//http.HandlerFunc handels requests to a specified web root
//Request count: GET/Counter
func handler(w http.ResponseWriter, r *http.Request){
    if r.URL.Path[1:] == "GET/Count"{
        fmt.Fprintf(w,"You are talking to instance %s:[internal_port]. This is the %dth request to this instance.", r.Host,counter )
        }else if r.URL.Path[1:] == "GET/temperature/"+appID+"/"+devID{
             a:= oTemp.OTemp(devID,appID,appAccessKey)
            fmt.Fprintf(w,"It's currently %.2f degrees in office %s/%s",a.PayloadFields["temperature"],appID,devID)
        }else if r.URL.Path[1:] == "favicon.ico"{
            counter++
        }else{
            fmt.Fprintf(w,"Nothing to see here. But the request still counts!")
        }
    }

    func main(){
        devID = os.Args[1]
        appID = os.Args[2]
        appAccessKey = os.Args[3]
    //http.HandleFunc("/", handler) tells the http package to handle all requests to the web root ("/") with handler
        http.HandleFunc("/", handler)
    //http.ListenAndServe, specifice that the http package should listen on port 8080
    //ListenAndServe always returns an error, since it only returns when an unexpected error occurs. 
    //In order to log that error we wrap the function call with log.Fatal.
        log.Fatal(http.ListenAndServe(":8080",nil))
    }


