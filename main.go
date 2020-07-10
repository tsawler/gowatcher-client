package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

var infoLog *log.Logger
var errorLog *log.Logger
var inProduction *bool

const gwcVersion = "1.0.0"

// App is the application config
type App struct {
	AllowFrom map[string]int
}

func main() {

	// create logs -- just writes to Stdout
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// TODO read these values from flags, .env, or whatever
	allowFrom := make(map[string]int)

	insecurePortPtr := flag.String("port", "", "Port")
	inProduction = flag.Bool("production", false, "application is in production")
	gwHost := flag.String("host", "", "goWatcher host IP")

	flag.Parse()

	insecurePort := *insecurePortPtr
	infoLog.Println("insecure port is", insecurePort)

	// always allow from localhost
	allowFrom["127.0.0.1"] = 1 // ipv4
	allowFrom["::1"] = 1       // ipv6
	allowFrom[*gwHost] = 1

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

	infoLog.Printf("Starting GoWatcher client v%s on port %s....", gwcVersion, insecurePort)

	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
