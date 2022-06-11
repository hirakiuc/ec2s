.DEFAULT_GOAL := default

.PHONY: build, deps, install, check, clean, default

build:
	go build

deps:
	go mod tidy
	go mod vendor

install:
	go install

check:
	golangci-lint run ./...

clean:
	go clean

default:
	make check
	make build
