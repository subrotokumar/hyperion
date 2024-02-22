# Simple Makefile for a Go project
DB=${DB_URL}

migrate-up:
	@migrate -path internal/db/migration -database ${DB} -verbose up

migrate-down:
	@migrate -path internal/db/migration -database "${DB}" -verbose down

proto:
	@protoc --proto_path=internal/proto --go_out=internal/pb --go_opt=paths=source_relative \
    --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
    internal/proto/*.proto

evans:
	@evans --host localhost --port 9090 -r repl

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
