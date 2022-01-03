package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.Handle("/", http.HandleFunc(home))
    mux.Handle("/snippet", http.HandleFunc(showSnippet))
    mux.Handle("/snippet/create", http.HandleFunc(createSnippet))

    // Serve the static files with relative path
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    // register the file server but strip /static prefix before the request
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Println("Serving up GO on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}