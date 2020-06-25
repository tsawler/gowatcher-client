#!/bin/bash
go build -o gwc *.go

./gwc -disk='/' -host='127.0.0.1' -port=':6001' -production=false