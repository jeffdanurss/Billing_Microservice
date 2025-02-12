# go official image
FROM golang:1.23-alpine AS builder

# work directory setting
WORKDIR /app

# files copy
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# app building
RUN CGO_ENABLED=0 GOOS=linux go build -o billing-microservice .

# latest image
FROM alpine:latest

# Instala dependencias necesarias
RUN apk add --no-cache bash

# Copia el binario desde el builder
WORKDIR /root/
COPY --from=builder /app/billing-microservice .

# Copia el archivo .env
COPY .env .

# Expone el puerto 8080
EXPOSE 8080

# Comando para iniciar la aplicación
CMD ["./billing-microservice"]