[![Go Report Card](https://goreportcard.com/badge/github.com/tsawler/gowatcher-client?style=flat-square)](https://goreportcard.com/report/github.com/tsawler/gowatcher-client) 


# goWatcher Client

The client for [goWatcher](https://github.com/tsawler/gowatcher)

Build:

~~~
env GOOS=linux GOARCH=amd64  go build -o gwc *.go
~~~

Run:

~~~bash
tcs@grendel gowatcher-client % ./gw -help
Usage of ./gw:
  -disk string
        disk to check (default "/")
  -host string
        goWatcher host IP
  -port string
        Port
  -production
        application is in production
~~~

Example:

~~~
./gw -host=178.128.231.48 -disk='/' -production=false -port=':6001'
~~~