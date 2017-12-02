package main

import (
    "log"
    "net/http"
    "./tournament"
)

func main(){
    // Add routes
    router := tournament.NewRouter()
    tournament.DataInitialization()
    // Initialize players and tournament
    log.Fatal(http.ListenAndServe(":8080", router))
}