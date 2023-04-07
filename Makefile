all: build test

build:
	go build -o server cmd/server/server.go

test:
	go test ./...