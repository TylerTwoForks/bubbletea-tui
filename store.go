package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	ID    int64
	Title string
	Body  string
}

type Store struct {
	conn *sql.DB
}

func (s *Store) Init() error {
	var err error
	s.conn, err = sql.Open("sqlite3", "./notes.db")
	if err != nil {
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		body TEXT NOT NULL
	);`
	if _, err := s.conn.Exec(createTableSQL); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetNotes() ([]Note, error) {
	rows, err := s.conn.Query("SELECT id, title, body FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Body); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func (s *Store) AddNote(note Note) error {
	if note.ID == 0 {
		note.ID = time.Now().UTC().UnixNano()
	}

	sql := `INSERT INTO notes (id, title, body) 
	VALUES (?, ?, ?) 
	ON CONFLICT(id) 
	DO UPDATE 
	SET title=excluded.title, body=excluded.body;`
	if _, err := s.conn.Exec(sql, note.ID, note.Title, note.Body); err != nil {
		return err
	}
	return nil
}
