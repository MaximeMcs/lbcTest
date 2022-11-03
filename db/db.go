package db

import (
	"database/sql"
	"fmt"
)

const (
	DB_HOST     = "host.docker.internal"
	DB_PORT     = "1234"
	DB_USER     = "root"
	DB_PASSWORD = "root"
	DB_NAME     = "lbc"
)

// DB set up
func SetupDB() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return nil, fmt.Errorf("db setup: %w", err)
	}

	return db, nil
}
