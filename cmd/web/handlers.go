package main

import(
    "errors"
    "fmt"
    "net/http"
    "strconv"

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
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }
    // create dummy data to be removed later
    title := "0 snail"
    content := "0 snaile\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
    expires := "7"

    // pass this data to the SnippetModel.Insert() method
    id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }
    // Redirect to the relevant page for the new snippet
    http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}