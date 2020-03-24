#!/usr/bin/make -f

test: version
	-go test -v

compile: version
	-go test -c -o ./parallel

testJSON: compile
	go test -json

version:
	@go version

separator:
	@echo
	@echo

all: test separator testJSON