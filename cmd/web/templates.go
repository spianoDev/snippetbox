package main

import "github.com/spianodev/snippetbox/pkg/models"

// This struct will allow any dynamic data to pass to the HTML templates
type templateData struct {
    Snippet *models.Snippet
    Snippets []*models.Snippet
}