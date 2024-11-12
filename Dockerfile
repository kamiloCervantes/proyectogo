# Imagen base de Go
FROM golang:1.20-alpine AS builder

# Configurar el directorio de trabajo
WORKDIR /app

# Copiar archivos de Go mod y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto de los archivos de la aplicación
COPY . .

# Compilar la aplicación
RUN go build -o myapp cmd/myapp/main.go

# Imagen final mínima para ejecutar la aplicación
FROM alpine:3.18

WORKDIR /app

# Copiar el binario compilado desde la imagen anterior
COPY --from=builder /app/myapp .

# Puerto en el que la aplicación se ejecuta
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./myapp"]
