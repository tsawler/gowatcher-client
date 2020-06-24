package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	// TODO read these values from flags, .env, or whatever
	insecurePort := ":4001"


	// create info log -- just writes to Stdout
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	// create http server
	srv := &http.Server{
		Addr:              insecurePort,
		ErrorLog:          errorLog,
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	infoLog.Printf("Starting HTTP server on port %s....", insecurePort)

	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
