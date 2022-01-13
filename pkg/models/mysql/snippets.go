package mysql

import (
  "database/sql"
  "github.com/spianodev/snippetbox/pkg/models"
)

// This struct defines a SnippetModel type that wraps a sql.DB connection pool
type SnippetModel struct {
  DB *sql.DB
}

// Function to insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
  return 0, nil
}

// Function to retrieve a specific snippet by id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
  return nil, nil
}

// Function to retrieve the latest 10 snippets
func (m, *SnippetModel) Latest()([]*models.Snippet, error) {
  return nil, nil
}