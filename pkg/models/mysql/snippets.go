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
  statement := `INSERT INTO snippets (title, content, created, expires)
  VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
  // Get the result by using the Exec() method with SQL statement, as first parameter
  result, err := m.DB.Exec(statement, title, content, expires)
  if err != nil {
   return 0, err
  }
  // Retrieve the int64 id of the newly inserted record
  id, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }
  // must convert the id result into an int
  return int(id), nil
}

// Function to retrieve a specific snippet by id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
  return nil, nil
}

// // Function to retrieve the latest 10 snippets
// func (m, *SnippetModel) Latest()([]*models.Snippet, error) {
//   return nil, nil
// }