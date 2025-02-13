# Etapa 1: Construcción del binario
FROM golang:1.23-alpine AS builder

# Configurar el directorio de trabajo
WORKDIR /app

# Copiar los archivos necesarios
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar el microservicio
RUN CGO_ENABLED=0 GOOS=linux go build -o billing-microservice .

# Etapa 2: Imagen final ligera
FROM alpine:latest

# Copiar el binario compilado
WORKDIR /root/
COPY --from=builder /app/billing-microservice .
# Copia el archivo .env (si es necesario)
COPY .env .

# Exponer el puerto 8080
EXPOSE 8080

# Comando para ejecutar el microservicio
CMD ["./billing-microservice"]