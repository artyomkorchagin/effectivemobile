package psqlsubscription

import (
	"database/sql"
	"fmt"

	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
	"github.com/pressly/goose/v3"
)

func migrationUp(db *sql.DB) error {
	migrationPath, err := getMigrationPath()
	if err != nil {
		return fmt.Errorf("[Goose up] %v", err)
	}
	if err := goose.SetDialect("pgx"); err != nil {
		return fmt.Errorf("[Goose up] Goose could not set the dialect to pgx: %v", err)
	}
	return goose.Up(db, migrationPath)
}

func migrationDown(db *sql.DB) error {
	migrationPath, err := getMigrationPath()
	if err != nil {
		return fmt.Errorf("[Goose down] %v", err)
	}
	if err := goose.SetDialect("pgx"); err != nil {
		return fmt.Errorf("[Goose down] Goose could not set the dialect to pgx: %v", err)
	}
	return goose.Up(db, migrationPath)
}

func getMigrationPath() (string, error) {
	migrationPath, err := helpers.GetProjectRoot()
	if err != nil {
		return "", fmt.Errorf("Migration path/root not found: %v", err)
	}

	migrationPath += "/migrations"
	return migrationPath, nil
}
