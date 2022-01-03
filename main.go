package main

import (
    "log"
    "net/http"
)

// Below is a home handler function which writes a byte slice containing "Hello from SpianoDev's Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
    // if statement to return 404 if the path doesn't match "/"
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Hello from SpianoDev's Snippetbox"))
}

// Below is a showSnippet handler function to display a specific showSnippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Here is that snippet you asked for..."))
}

// Below is a createSnippet handler function that will make a new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
    // if statement to return 405 if the method is not POST
    if r.Method != http.MethodPost {
        // must call w.WriteHeader before w.Write to send anything other than 200. This can only be used once per response
        w.WriteHeader(405)
        w.Write([]byte("Method NOT Allowed. Use POST."))
        return
    }
    w.Write([]byte("Creating a new snippet right away!"))
}

// Below is the main function that
// First initializes a new servemux and registers the home function as "/" in the URL pattern
// Second starts a new web server by passing in the TCP network address and the new servemux
// Third log and throw an error and exit

    // can also use the DefaultServeMux global variable for identical results without declaring mux: http.HandleFunc("/", home)
    // [This is NOT recommended for prod apps because of the global variable that allows any package to access it]

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    log.Println("Serving up GO on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}

