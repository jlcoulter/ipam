# Go API Service — Makefile

BINARY   := bin/server
PORT     ?= 8080

.PHONY: run build test vet lint docker clean

run: build
	./$(BINARY)

build:
	go build -o $(BINARY) ./cmd/server

test:
	go test -race -cover ./...

vet:
	go vet ./...

lint: vet
	golangci-lint run ./...

docker:
	docker build -t go-api-template .

clean:
	rm -rf bin/