package main

import (
	"database/sql"

	_ "github.com/jackc/pgconn" // need this and next two for pgx
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func checkPostgres(dsn string) (bool, string, string, error) {

	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return false, "Cannot connect to database!", "", err
	}

	err = d.Ping()
	if err != nil {
		return false, "Cannot ping to database!", "", err
	}

	return true, "Pinged database successfully!", "", nil
}
