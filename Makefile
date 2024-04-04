BINARY_NAME=rempass

.PHONY: build
build:
	@echo 'Building for Linux'
	go build -o=./bin/${BINARY_NAME} ./cmd/cli

.PHONY: start
start:
	go run ./cmd/cli
