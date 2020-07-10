package main

import (
	"database/sql"

	_ "github.com/jackc/pgconn" // need this and next two for pgx
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func checkPostgres(dsn string) (bool, string, string, int, error) {

	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return false, "Cannot connect to database!", "", 2, err
	}
	defer d.Close()

	err = d.Ping()
	if err != nil {
		return false, "Cannot ping to database!", "", 2, err
	}

	return true, "Pinged database successfully!", "", 1, nil
}
