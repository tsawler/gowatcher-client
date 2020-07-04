[![Go Report Card](https://goreportcard.com/badge/github.com/tsawler/gowatcher-client?style=flat-square)](https://goreportcard.com/report/github.com/tsawler/gowatcher-client) 


# goWatcher Client

The client for [goWatcher](https://github.com/tsawler/gowatcher)

Build:

~~~
env GOOS=linux GOARCH=amd64  go build -o gwc *.go
~~~

Run:

~~~
tcs@grendel gowatcher-client % ./gwc -help
Usage of ./gwc:
  -host string
        goWatcher host IP
  -port string
        Port
  -production
        application is in production
~~~

Example:

~~~
./gwc -host=178.128.231.48 -production=false -port=':6001'
~~~
