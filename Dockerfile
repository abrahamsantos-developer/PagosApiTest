# 1. Usa una imagen base ligera de Go
FROM golang:1.22-alpine

# 2. Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# 3. Copiar los archivos de Go mod para descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# 4. Copiar el resto de los archivos del proyecto
COPY . .

# 5. Compilar la aplicación
RUN go build -o app ./cmd/api

# 6. Exponer el puerto en el que correrá la API
EXPOSE 3000

# 7. Comando para ejecutar la aplicación cuando el contenedor esté corriendo
CMD ["./app"]
