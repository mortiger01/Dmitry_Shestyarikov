# Этап 1: Билдер (Go 1.24)
FROM golang:1.25-alpine AS builder

# Устанавливаем git и swag
RUN apk add --no-cache git
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

# Копируем go.mod и go.sum для кеширования
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код
COPY . .

# Генерируем swagger-документацию
RUN swag init -g cmd/main.go -o docs

# Собираем статический бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o pipe-api ./cmd/main.go

# Этап 2: Финальный образ (минимальный alpine)
FROM alpine:latest

# Устанавливаем CA-сертификаты (нужны для OAuth и HTTPS)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник и папки из билдера
COPY --from=builder /app/pipe-api .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/docs ./docs

EXPOSE 4200

CMD ["./pipe-api"]