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

## Running with Caddy

To run as part of existing site, with Caddy 2.x, use this redirect (assuming url is https://www.goblender.ca/gowatcher/status,
and gwc is running on port 6001):

~~~
www.goblender.ca, goblender.ca {
	encode zstd gzip
	import static
	import security

	log {
		output file /var/www/go.verilion.com/logs/caddy-access.log
		format single_field common_log
	}

	# for tus
	reverse_proxy /files/* localhost:1080

	# for gowatcher client
	route /gowatcher/status/* {
		uri strip_prefix /gowatcher/status
		reverse_proxy localhost:6001
	}

	# for goblender
	reverse_proxy  http://localhost:7004
}
~~~