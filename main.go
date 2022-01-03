package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
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
    // Get value of id parameter, convert to int, unless values is <1, then return 404
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }
    /* Replace original message with formatted function that interpolates the id value
    w.Write([]byte("Here is that snippet you asked for..."))
    */
    fmt.Fprintf(w, "Serving Snippet ID %d", id)
}

// Below is a createSnippet handler function that will make a new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
    // if statement to return 405 if the method is not POST
    if r.Method != http.MethodPost {
        // use allow header method with header name and header value as parameters
        w.Header().Set("Allow", http.MethodPost)
/*
//         // must call w.WriteHeader before w.Write to send anything other than 200. This can only be used once per response
//         w.WriteHeader(405)
//         w.Write([]byte("Method NOT Allowed."))
//        // This entire block can be replaced with an error method
*/
        http.Error(w, "Method NOT Allowed", 405)
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

// setting cache-control
/* This would overwrite every existing header
w.Header().Set("Cache-Control", "public, max-age=31536000")
*/

/* By contrast this appends to an existing cache and can be called multiple times
w.Header().Add("Cache-Control", "public")
w.Header().Add("Cache-Control", "max-age=331536000")
*/

/* This command deletes all values in the "Cache-Control" header
w.Header().Del("Cache-Control")

This command retrieves the first value for the "Cache-Control" header
w.Header().Get("Cache-Control")
*/

// The manual process for sending JSON responses
/*
w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{"name": "spianoDev"}`))
*/
