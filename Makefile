.PHONY: build test run clean

BINARY_NAME=cube-cli

build:
	go build -o $(BINARY_NAME) ./cmd/cli/

test:
	go test ./internal/... ./cmd/cli/... ./tests/...

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
