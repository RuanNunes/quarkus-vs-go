# Etapa 1: Build
FROM golang:1.23.3 as builder

# Define o diretório de trabalho no container
WORKDIR /app

# Copia os arquivos necessários para o build
COPY . .
RUN go mod download



# Compila o binário de forma completamente estática
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-extldflags '-static'" -o go-api .

# Etapa 2: Execução (Distroless)
FROM gcr.io/distroless/base-debian11

# Define o diretório de trabalho no container
WORKDIR /app

# Copia o binário da etapa de build
COPY --from=builder /app/go-api .

# Expõe a porta usada pelo aplicativo
EXPOSE 8090

# Comando para iniciar a aplicação
CMD ["/app/go-api"]