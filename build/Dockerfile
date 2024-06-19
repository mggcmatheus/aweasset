# Estágio de construção
FROM golang:alpine AS builder

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie o código fonte para o container
COPY . .

# Compile o aplicativo
RUN go build -ldflags="-s -w" -o app cmd/main.go

# Estágio de produção
FROM alpine:latest

# Copie o arquivo de aplicação compilado do estágio de construção
COPY --from=builder /app/app /app/app

# Copie os arquivos de configuração necessários para o contêiner
COPY .env /app/.env
COPY config-docker.yml /app/config-docker.yml

# Defina o diretório de trabalho para o aplicativo
WORKDIR /app

# Copie o arquivo start.sh para o contêiner
COPY build/scripts/start.sh /app/start.sh
RUN chmod +x /app/start.sh

# Defina o comando padrão para iniciar o aplicativo
CMD ["/app/start.sh"]