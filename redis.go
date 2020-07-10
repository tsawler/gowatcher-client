package main

import (
	"github.com/mediocregopher/radix/v3"
)

func checkRedis(dsn string) (bool, string, string, int, error) {

	pool, err := radix.NewPool("tcp", dsn, 1)
	if err != nil {
		return false, "Error connecting to redis", "", 2, err
	}
	defer pool.Close()

	return true, "Pinged redis successfully!", "", 1, nil
}
