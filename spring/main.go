package main

import (
    "fmt"
    "log"
    "net/http"   
    "github.com/gorilla/mux" 
)

func main() {
    fmt.Println("started-service")

    r := mux.NewRouter()

    // OPTIONS
    // front end is aws.com, back end is google.com
    // options could match for different domain
    // different company or team
    // cross domain request
    r.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST", "OPTIONS")
    r.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET", "OPTIONS")
    log.Fatal(http.ListenAndServe(":8080", r))
}
