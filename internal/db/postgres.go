package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func New(connStr string) (*DB, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err = conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to connect to conn: %w", err)
	}

	createTableSQL := `
		CREATE TABLE IF NOT EXISTS tasks (
				id SERIAL PRIMARY KEY,
				title TEXT NOT NULL,
				category TEXT,
				description TEXT,
				amount INTEGER,
				start_date TIMESTAMPTZ,
				end_date TIMESTAMPTZ,
				notification_date TIMESTAMPTZ,
				created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
				updated_at TIMESTAMPTZ
		);
	`

	if _, err := conn.Exec(createTableSQL); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create tasks table: %w", err)
	}

	return &DB{Conn: conn}, nil
}
