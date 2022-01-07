package main

import (
    "flag"
    "log"
    "net/http"
)

func main() {
    // Adding a command line flag
    addr := flag.String("addr", ":4000", "HTTP network address")
    // Function below parses the command line flag, reads it and assigns the addr variable
    // The returned value is a pointer to the flag value so it needs to be prefixed with '*'
    flag.Parse()

    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    // Serve the static files with relative path
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    // register the file server but strip /static prefix before the request
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Println("Serving up GO on %s", *addr)
    err := http.ListenAndServe(*addr, mux)
    log.Fatal(err)
}