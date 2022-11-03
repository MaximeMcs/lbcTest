package db

import (
	"database/sql"
	"fmt"
	"root/utils"
)

const (
	DB_HOST     = "host.docker.internal"
	DB_PORT     = "1234"
	DB_USER     = "root"
	DB_PASSWORD = "root"
	DB_NAME     = "lbc"
)

// DB set up
func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	utils.CheckErr(err)

	return db
}
