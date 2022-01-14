package main

import (
    "net/http"

    "github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
    // Leverage a middleware chain for standard requests
    standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    mux := http.NewServeMux()
    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet", app.showSnippet)
    mux.HandleFunc("/snippet/create", app.createSnippet)

    // Serve the static files with relative path
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    // register the file server but strip /static prefix before the request
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    return standardMiddleware.Then(mux)
}