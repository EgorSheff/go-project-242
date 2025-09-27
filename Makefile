BIN_PATH = bin/hexlet-path-size

build:
	go build -ldflags="-w -s" -gcflags=all="-l -B" -o $(BIN_PATH) ./cmd/hexlet-path-size

lint:
	golangci-lint run ./...

test:
	go test -v

clean:
	rm bin/* || true # Ignore errors
