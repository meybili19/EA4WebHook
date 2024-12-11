# Etapa 1: Compilación
FROM golang:1.23.4 AS builder

WORKDIR /app

# Copiar archivos del proyecto
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Generar la documentación Swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init

# Compilar la aplicación
RUN go build -o main .

# Etapa 2: Imagen final con GLIBC preinstalado
FROM debian:bookworm-slim

WORKDIR /app

# Instalar dependencias mínimas (si es necesario)
RUN apt-get update && apt-get install -y \
    ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Copiar los binarios y la documentación desde la etapa de compilación
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Exponer el puerto de la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]

