package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var infoLog *log.Logger
var errorLog *log.Logger
var diskToCheck string
var inProduction bool

// DiskThreshold is the warning threshold for disks
const DiskThreshold = 90

//MemoryThreshold is hte warning threshold for memory
const MemoryThreshold = 80

// App is the application config
type App struct {
	AllowFrom map[string]int
}

func main() {
	// TODO read these values from flags, .env, or whatever
	inProduction = true
	insecurePort := ":4001"
	allowFrom := make(map[string]int)

	// always allow from localhost
	allowFrom["127.0.0.1"] = 1 // ipv4
	allowFrom["::1"] = 1       // ipv6

	diskToCheck = "/"

	// create logs -- just writes to Stdout
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := App{
		AllowFrom: allowFrom,
	}

	// create http server
	srv := &http.Server{
		Addr:              insecurePort,
		ErrorLog:          errorLog,
		Handler:           routes(app),
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
