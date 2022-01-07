package main

import(
    "fmt"
    "html/template"
    "net/http"
    "strconv"
)

func (app *application)home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }
    // adding the templates with home first
    files := []string{
        "./ui/html/home.page.tmpl",
        "./ui/html/base.layout.tmpl",
        "./ui/html/footer.partial.tmpl",
    }
    // adding the files defined above
    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }
    // Write the template content
    err = ts.Execute(w, nil)
    if err != nil {
        app.serverError(w, err)
    }
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }

    fmt.Fprintf(w, "Serving Snippet ID %d", id)
}

func (app *application)createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

    w.Write([]byte("Creating a new snippet right away!"))
}