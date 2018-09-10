package main

import (
"fmt"
"log"
"net/http"
)

/*
The task states that the number of requests is to be counted and printed out on request (GET/Counter).
However, for every request by the handler the browser will also request the website icon ("favicon.ico").
Consequently, a request by the user will always result in two requests to the server.
Since most users are not aware of this, the counter will only count one of these two requests.
In case that both requests are to be counted move the "counter++" outside of the if-else statement and omit the else-if case
*/


//Counts the Requests to the server 
var counter int = 1



//http.HandlerFunc handels requests to a specified web root
//Request count: GET/Counter
func handler(w http.ResponseWriter, r *http.Request){
    if r.URL.Path[1:] == "GET/Count"{
        fmt.Fprintf(w,"You are talking to instance %s:[internal_port]. This is the %dth request to this instance.", r.Host,counter )
    }else if r.URL.Path[1:] == "favicon.ico"{
        counter++
    }else{
        fmt.Fprintf(w,"Nothing to see here. But the request still counts!")
    }
}

func main(){
    //http.HandleFunc("/", handler) tells the http package to handle all requests to the web root ("/") with handler
    http.HandleFunc("/", handler)
    //http.ListenAndServe, specifice that the http package should listen on port 8080
    //ListenAndServe always returns an error, since it only returns when an unexpected error occurs. 
    //In order to log that error we wrap the function call with log.Fatal.
    log.Fatal(http.ListenAndServe(":8080",nil))
}


