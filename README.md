# parallel

This package runs tests and sub-tests **both in parallel**. The purpose of
this is to illustrate a bug in the parsing of test output in GoLand 2019.3.3
(Build #GO-193.6494.61) when used to run these tests using Go 1.14.

At SmartyStreets we have built [our own testing library](https://github.com/smartystreets/gunit)
which makes heavy use of parallel tests and sub-tests (see [`t.Parallel()`](https://golang.org/pkg/testing/#T.Parallel)).
After upgrading to Go 1.14 we noticed that failure output wasn't being presented
correctly by the GoLand test runner.
