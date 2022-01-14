package main

import (
    "html/template"
    "path/filepath"

    "github.com/spianodev/snippetbox/pkg/models"
)
// This struct will allow any dynamic data to pass to the HTML templates
type templateData struct {
    Snippet *models.Snippet
    Snippets []*models.Snippet
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

        ts, err := template.ParseFiles(page)
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