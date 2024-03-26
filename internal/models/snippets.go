package models

import (
	"database/sql"
	"time"
)

// snippet type that will hold the data for a single
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// function to insert a new snippet ino the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// sql statement to insert a new record into the snippets table
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// use the Exec() method on the embedded connection pool to execute the statement
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// function to get a specific snippet based on its ID
//func (m *SnippetModel) Get(id int) (Snippet, error) {
//	return nil, nil
//}

// function to get the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
