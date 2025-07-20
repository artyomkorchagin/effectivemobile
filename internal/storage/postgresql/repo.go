package psqlsubscription

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) (*Repository, error) {
	err := migrationUp(db)

	if err != nil {
		return nil, fmt.Errorf("Could not migrate database: %w", err)
	}

	return &Repository{
		db: db,
	}, nil
}
