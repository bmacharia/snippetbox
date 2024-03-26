package models

import (
	"database/sql"
	"errors"
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
func (m *SnippetModel) Get(id int) (Snippet, error) {

	// write the sql statement that we want to execute
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`
	// use the QueryRow() method on the connection pool to execute our to execute our SQL statement
	// this returns a pointer to a sql.Row object which holds the result from the database
	row := m.DB.QueryRow(stmt, id)

	// z new zeroed Snippet struct
	var s Snippet

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// if the query returns no rows from the database, then row.Scan() will return a sql.ErrNoRows error
		// that error is handle by the errors.Is() function
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}
	// if everything worked as expected, then return the Snippet data as normal
	return s, nil
}

// function to get the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	// use the Query() method on the connection pool to execute our SQL statement
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// we defer rows.Close() to ensure that the result set is always properly closed before the Latest() method returns
	defer rows.Close()

	// a slice of Snippet struct to hold the data
	var snippets []Snippet

	for rows.Next() {
		var s Snippet

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
