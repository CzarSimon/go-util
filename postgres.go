package main

import (
	"database/sql"
	"fmt"
)

//ConnectPG creates a connection to a postgresql DB
func ConnectPG(config PGConfig) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DB)
	db, err := sql.Open("postgres", connStr)
	CheckErrFatal(err)
	err = db.Ping()
	CheckErrFatal(err)
	return db
}
