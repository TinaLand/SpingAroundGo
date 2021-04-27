package main

import (
    "fmt"
    "log"
    "net/http"   
    "github.com/gorilla/mux" 
    jwtmiddleware "github.com/auth0/go-jwt-middleware"
    jwt "github.com/form3tech-oss/jwt-go"
)

func main() {
    fmt.Println("started-service")

    // OPTIONS
    // front end is aws.com, back end is google.com
    // options could match for different domain
    // different company or team
    // cross domain request

    jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
            return []byte(mySigningKey), nil
        },
        SigningMethod: jwt.SigningMethodHS256,
    })

    r := mux.NewRouter()

    r.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST", "OPTIONS")
    r.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET", "OPTIONS")
    r.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST", "OPTIONS")
    r.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST", "OPTIONS")
    r.Handle("/post/{id}", jwtMiddleware.Handler(http.HandlerFunc(deleteHandler))).Methods("DELETE", "OPTIONS")

    log.Fatal(http.ListenAndServe(":8080", r))
}
