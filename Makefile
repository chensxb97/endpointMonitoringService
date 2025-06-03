APP_NAME := endpointMonitoringService
SRC := ./cmd/main.go
CONFIG := ./config.yaml

# Commands
run:
	@echo "Running $(APP_NAME)..."
	go run $(SRC) $(CONFIG)

build:
	@echo "Building $(APP_NAME)..."
	go build -o $(APP_NAME) $(SRC)

clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

test:
	@echo "Running tests..."
	go test ./...

lint:
	@echo "Linting code..."
	golangci-lint run

check-config:
	@echo "Validating YAML config..."
	yq e . $(CONFIG)

help:
	@echo "Available commands:"
	@echo "  run          - Run the app"
	@echo "  build        - Build the app"
	@echo "  clean        - Remove built binaries"
	@echo "  test         - Run tests"
	@echo "  lint         - Run linter"
	@echo "  check-config - Validate YAML config"