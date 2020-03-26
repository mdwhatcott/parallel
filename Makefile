#!/usr/bin/make -f

test: version
	-go test -v

test2JSONCleanup: test2JSON
	rm ./parallel

test2JSON: version compileTest
	-go tool test2json -t ./parallel -test.v

compileTest:
	go test -c -o ./parallel

testJSON:
	-go test -json

version:
	@go version

all: test testJSON test2JSONCleanup