package database

import (
	"database/sql"
	"fmt"
	"gator/internal/config"
)

type DB struct {
	SQL     *sql.DB
	Queries *Queries
}

func ConnectToDB(cfg config.Config) (*DB, error) {
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return &DB{SQL: db, Queries: New(db)}, nil
}
