package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Storage struct {
	DB *sql.DB
}

func New(storagePath string) (Storage, error) {
	const fn = "database.New"

	db, err := sql.Open("sqlite", storagePath)
	if err != nil {
		return Storage{}, fmt.Errorf("%s: %w", fn, err)
	}

	stmt, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS person(
      id INTEGER PRIMARY KEY,
      email TEXT UNIQUE,
      phone TEXT UNIQUE,
      firstName TEXT UNIQUE NOT NULL,
      lastName TEXT UNIQUE NOT NULL);
    CREATE INDEX IF NOT EXISTS person_firstName ON person(firstName);
    CREATE INDEX IF NOT EXISTS person_lastName ON person(lastName);
  `)
	if err != nil {
		return Storage{}, fmt.Errorf("%s: %w", fn, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return Storage{}, fmt.Errorf("%s: %w", fn, err)
	}

	return Storage{DB: db}, nil
}
