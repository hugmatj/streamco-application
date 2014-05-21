#!/bin/sh

echo "...fetching dependencies"
go get -v

echo "running webserver"
echo "=================\n"
go run server.go
