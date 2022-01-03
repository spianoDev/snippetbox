package main

import (
    "log"
    "net/http"
)

// Below is a home handler function which writes a byte slice containing "Hello from SpianoDev's Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from SpianoDev's Snippetbox"))
}

// Below is the main function that
// First initializes a new servemux and registers the home function as "/" in the URL pattern
// Second starts a new web server by passing in the TCP network address and the new servemux
// Third log and throw an error and exit

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    log.Println("Serving up GO on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}

