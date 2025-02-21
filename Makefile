default: fmt lint install

build:
	go build -v ./...

install: build
	go install -v ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...

.PHONY: fmt lint build install
