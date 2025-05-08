APP_NAME := parse_cbr
PKG := ./cmd
.PHONY: all build run clean fmt vet help

all: build

# Build executable file
build:
	go mod tidy
	go build -o $(APP_NAME) $(PKG)

# Run builded file
run: build
	./$(APP_NAME)

# Check code using 'vet'
vet:
	go vet ./...

# Format code using 'go fmt'
fmt:
	go fmt ./...

# Delete built file
clean:
	-del /F /Q $(APP_NAME)

# Print all commands
help:
	@echo "Usage: make <command>"
	@echo Available commands:
	@echo   all   - Build the project (default)
	@echo   build - Compile the Go application
	@echo   run   - Build and run the application
	@echo   vet   - Run Go vet (static analysis)
	@echo   fmt   - Run Go fmt (code formatting)
	@echo   clean - Remove built files
