package main

import (
    "bytes"
    "fmt"
    "net/http"
    "runtime/debug"
)

// Cache template helper to render templates from the cache
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
    // retrieve the appropriate template from cache based on name
    ts, ok := app.templateCache[name]
    if !ok {
        app.serverError(w, fmt.Errorf("Snippets cannot find the template you are looking for: %s.", name))
        return
    }

    // Initialize a buffer
    buffer := new(bytes.Buffer)

    // Write to buffer first so errors return correctly to user
    err := ts.Execute(buffer, td)
    if err != nil {
        app.serverError(w, err)
        return
    }

    buffer.WriteTo(w)
}

// Error helper writes message and sends 500 Internal Server Error to user
func (app *application) serverError(w http.ResponseWriter, err error) {
    trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
    app.errorLog.Output(2, trace)

    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// ClientError helper sends specific status
func (app *application) clientError(w http.ResponseWriter, status int) {
    http.Error(w, http.StatusText(status), status)
}

// NotFound helper sends 404 Error
func (app *application) notFound(w http.ResponseWriter) {
    app.clientError(w, http.StatusNotFound)
}