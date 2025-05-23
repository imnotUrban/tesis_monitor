# Usa una imagen de Go ligera
FROM golang:1.21-alpine

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos necesarios
COPY . .

# Compila el binario
RUN go build -o monitor .

# Define el entrypoint
CMD ["./monitor"]
