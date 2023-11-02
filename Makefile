.PHONY: build lint test tidy vet

default: all

all: tidy vet lint test build

build:
	go build ./...

lint:
	golangci-lint run --config .golangci.yaml

test:
	go test ./...

tidy:
	go mod tidy

vet:
	go vet ./...
