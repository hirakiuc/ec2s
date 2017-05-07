.DEFAULT_GOAL := default

build:
	go build

install:
	go install

check:
	go vet
	golint

clean:
	go clean

default:
	make check
	make build
