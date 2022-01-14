package main

import(
    "errors"
    "fmt"
    "net/http"
    "strconv"
    "strings"
    "unicode/utf8"

    "github.com/spianodev/snippetbox/pkg/models"
)

func (app *application)home(w http.ResponseWriter, r *http.Request) {

    s, err := app.snippets.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }

    app.render(w, r, "home.page.tmpl", &templateData{
        Snippets: s,
    })
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get(":id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }

    s, err := app.snippets.Get(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }

    app.render(w, r, "show.page.tmpl", &templateData{
        Snippet: s,
    })
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
    app.render(w, r, "create.page.tmpl", nil)
}

func (app *application)createSnippet(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    title := r.PostForm.Get("title")
    content := r.PostForm.Get("content")
    expires := r.PostForm.Get("expires")

    errors := make(map[string]string)

    if strings.TrimSpace(title) == "" {
        errors["title"] = "ERROR, PLEASE ADD A TITLE!"
    } else if utf8.RuneCountInString(title) > 160 {
        errors["title"] = "Title is too long (maximum is 160 characters)"
    }

    if strings.TrimSpace(content) == "" {
        errors["content"] = "ERROR, PLEASE ADD CONTENT!"
    }

    if strings.TrimSpace(expires) == "" {
        errors["expires"] = "ERROR, PLEASE SELECT AN EXPIRATION!"
    } else if expires != "365" && expires != "7" && expires != "1" {
        errors["expires"] = "Invalid field"
    }

    // Dump any errors into plain text HTTP response and return to user
    if len(errors) > 0 {
        fmt.Fprint(w, errors)
        return
    }
    
    // pass the form data to the SnippetModel.Insert() method
    id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }
    // Redirect to the relevant page for the new snippet
    http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}