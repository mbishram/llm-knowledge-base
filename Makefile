.PHONY: build test run clean

BINARY_NAME=llm-knowledge-base
BIN_DIR=bin

build:
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/llm-knowledge-base

test:
	go test ./internal/...

run: build
	./$(BIN_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BIN_DIR)
	go clean
