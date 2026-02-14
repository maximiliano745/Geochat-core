FROM golang:1.24-alpine

# Evitamos toolchains externos
ENV GOTOOLCHAIN=local

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# 1. Copiamos solo los archivos de módulos para descargar dependencias
COPY backend/go.mod backend/go.sum* ./

# 2. Hack inicial para poder bajar las librerías
RUN sed -i 's/go 1\..*/go 1.24.0/' go.mod && go mod download

# 3. Copiamos todo el resto del código del backend
COPY backend/ .

# 4. HACK FINAL: Volvemos a forzar la versión por si el COPY la sobreescribió
# Y compilamos en la misma línea para que no haya dudas
RUN sed -i 's/go 1\..*/go 1.24.0/' go.mod && \
    go build -o main ./cmd/server/main.go

EXPOSE 8080
CMD ["./main"]