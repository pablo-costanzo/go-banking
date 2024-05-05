# Usa una imagen específica en vez de latest para asegurar compatibilidad.
FROM golang:1.22.1

# Establece el directorio de trabajo dentro del contenedor.
WORKDIR /app

# Copia el archivo go.mod y go.sum para manejar las dependencias.
COPY go.mod go.sum ./

# Descarga las dependencias en caché si los archivos go.mod y go.sum no cambian.
RUN go mod download

# Copia el resto del código fuente del proyecto al contenedor.
COPY . .

# Compila la aplicación para reducir tiempos de arranque.
RUN go build -o /go-banking

# El comando que se ejecutará al iniciar el contenedor.
CMD ["/go-banking"]
