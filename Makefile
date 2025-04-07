.PHONY: build test lint fmt run

build:
	go build -v ./...

test:
	go test -v ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...

run-chat-server:
	go run ./cmd/chat-server/main.go
