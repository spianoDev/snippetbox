package mysql

import (
  "database/sql"
  "errors"
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
  statement := `SELECT id, title, content, created, expires FROM snippets
  WHERE expires > UTC_TIMESTAMP() AND id = ?`

  // Get a pointer to the row that matches the above statement
  row := m.DB.QueryRow(statement, id)
  // New variable for entering the values for finding the desired row
  s := &models.Snippet{}
  // look for the values in each row leveraging the 's' variable
  err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, models.ErrNoRecord
    } else {
      return nil, err
    }
  }

  return s, nil
}

// // Function to retrieve the latest 10 snippets
func (m *SnippetModel) Latest()([]*models.Snippet, error) {
  statement := `SELECT id, title, content, created, expires FROM snippets
  WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

  // in this case we simply need to query the above statement
  rows, err := m.DB.Query(statement)
  if err != nil {
    return nil, err
  }
  // close before returning but after checking for errors
  defer rows.Close()
  // variable to hold multiple snippets
  snippets := []*models.Snippet{}
  // iterate over the results
  for rows.Next() {
    s := &models.Snippet{}
    err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err != nil {
      return nil, err
    }
    snippets = append(snippets, s)
  }
  // look for any errors that happened during the loop
  if err = rows.Err(); err != nil {
    return nil, err
  }

  return snippets, nil
}