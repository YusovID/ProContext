APP_NAME := robot_race.exe
PKG := ./cmd
.PHONY: all build run clean fmt vet help

all: build

# Сборка исполняемого файла
build:
	go build -o $(APP_NAME) $(PKG)

# Запуск скомпилированного файла
run: build
	./$(APP_NAME)

# Проверка кода стандартным инструментом vet
vet:
	go vet ./...

# Форматирование кода стандартным инструментом go fmt
fmt:
	go fmt ./...

# Удаление скомпилированного файла 
clean:
	-del /F /Q $(APP_NAME)

# Вывод списка доступных команд
help:
	@echo "Usage: make <command>"
	@echo Available commands:
	@echo   all       - Build the project (default)
	@echo   build     - Compile the Go application
	@echo   run       - Build and run the application
	@echo   vet       - Run Go vet (static analysis)
	@echo   fmt       - Run Go fmt (code formatting)
	@echo   clean     - Remove built files
