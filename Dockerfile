# Etapa de construção
FROM golang:1.23.1-alpine as builder

# Configura o diretório de trabalho
WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -o executavel

# Etapa final
FROM scratch

# Copia o binário da etapa de construção
COPY --from=builder /app/executavel /

# Define o ponto de entrada
ENTRYPOINT ["/executavel"]
