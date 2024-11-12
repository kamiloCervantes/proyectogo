# Variables
APP_NAME = myapp
CMD_DIR = ./cmd/$(APP_NAME)
BUILD_DIR = ./bin

# Comandos
.PHONY: all build run test clean deps fmt lint help

all: build

build:
	@echo "Compilando $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

run: build
	@echo "Ejecutando $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "Ejecutando pruebas..."
	@go test ./...

clean:
	@echo "Limpiando..."
	@rm -rf $(BUILD_DIR)

deps:
	@echo "Descargando dependencias..."
	@go mod download

fmt:
	@echo "Formateando código..."
	@go fmt ./...

lint:
	@echo "Analizando código..."
	@golangci-lint run

help:
	@echo "Comandos disponibles:"
	@echo "  make build      - Compila la aplicación"
	@echo "  make run        - Compila y ejecuta la aplicación"
	@echo "  make test       - Ejecuta pruebas"
	@echo "  make clean      - Limpia archivos de compilación"
	@echo "  make deps       - Descarga dependencias"
	@echo "  make fmt        - Formatea el código"
	@echo "  make lint       - Analiza el código"
	@echo "  make help       - Muestra esta ayuda"
