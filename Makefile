clean:
	go clean

dependencies:
	go get
	go get github.com/smartystreets/goconvey

build: clean dependencies
	go build aoc.go

test: build
	go test ./...

ci: clean dependencies build test

default: build

.PHONY: test
