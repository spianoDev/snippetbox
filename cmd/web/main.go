package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    log.Println("Serving up GO on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}