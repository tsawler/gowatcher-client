package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func checkMariaDB(dsn string) (bool, string, string, error) {
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		errorLog.Println(err)
		return false, "Cannot connect to database!", "", err
	}

	err = d.Ping()
	if err != nil {
		errorLog.Println(err)
		return false, "Cannot ping to database!", "", err
	}

	return true, "Pinged database successfully!", "", nil
}
