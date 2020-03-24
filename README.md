# parallel

This package runs tests and sub-tests **both in parallel**. The purpose of
this is to illustrate a bug in the output of `go test -json` (and also
`go tool test2json`) using Go 1.14 and Go 1.14.1. See lines 77-78 of 
"output-go1.14.1.txt" for evidence of this. 

At SmartyStreets we have built [our own testing library](https://github.com/smartystreets/gunit)
which makes heavy use of parallel tests and sub-tests (see [`t.Parallel()`](https://golang.org/pkg/testing/#T.Parallel)).
After upgrading to Go 1.14 we noticed that failure output wasn't being presented
correctly by the GoLand test runner.

