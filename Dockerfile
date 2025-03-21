# Use a imagem base do Golang
FROM golang:1.20-alpine AS builder

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie go.mod e go.sum (se existirem)
COPY go.mod go.sum* ./

# Baixe as dependências
RUN go mod download

# Copie o restante do código da aplicação
COPY . .

# Compile a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o user-api .

# Use uma imagem mínima para a execução
FROM alpine:latest

# Instale certificados para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copie o binário compilado da etapa anterior
COPY --from=builder /app/user-api .

# Exponha a porta que a aplicação usa
EXPOSE 3000

# Comando para iniciar a aplicação
CMD ["./user-api"]