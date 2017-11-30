package tournament

import (
    "log"
    "net/http"
)

func main(){
    // Add routes
    router := NewRouter()
    // Initialize players and tournamnet
    log.Fatal(http.ListenAndServe(":8080", router))
    datainit()
}