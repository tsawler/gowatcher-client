package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func checkMariaDB(dsn string) (bool, string, string, int, error) {
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		errorLog.Println(err)
		return false, "Cannot connect to database!", "", 2, err
	}
	defer d.Close()

	err = d.Ping()
	if err != nil {
		errorLog.Println(err)
		return false, "Cannot ping to database!", "", 2, err
	}

	return true, "Pinged database successfully!", "", 1, nil
}
