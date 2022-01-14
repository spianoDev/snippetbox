package main

import (
    "html/template"
    "path/filepath"
    "time"

    "github.com/spianodev/snippetbox/pkg/models"
)
// This struct will allow any dynamic data to pass to the HTML templates
type templateData struct {
    CurrentYear int
    Snippet *models.Snippet
    Snippets []*models.Snippet
}

// Function to make date human friendly
func humanDate(t time.Time) string {
    return t.Format("Jan 02, 2006 at 13:05 MST")
}

// Make a global variable that passes the human friendly date
var functions = template.FuncMap{
    "humanDate": humanDate,
}

// Function to cache pages
func newTemplateCache(dir string) (map[string]*template.Template, error) {
    cache := map[string]*template.Template{}
    // filepath.Glob method gets a slice of all paths with extension .page.tmpl
    pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
    if err != nil {
        return nil, err
    }

    for _, page := range pages {
        // assign each file name in the loop to the name variable
        name := filepath.Base(page)

        ts, err := template.New(name).Funcs(functions).ParseFiles(page)
        if err != nil {
            return nil, err
        }

        // parse any additional layout templates to the template set
        ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
        if err != nil {
                return nil, err
        }

        // parse any partial templates (like the footer)
        ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
        if err != nil {
                return nil, err
        }

        cache[name] = ts
    }

    return cache, nil
}