# Этап сборки
FROM golang:1.25 AS build

WORKDIR /app

# Копируем исходники сервиса
COPY auth/ ./

# Копируем общие пакеты
COPY lib/ ../lib/

# Копируем go.mod и go.sum сервиса
COPY auth/go.mod auth/go.sum ./
RUN go mod download

# Собираем бинарный файл нашего приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/auth auth/cmd/

# Добавляем пользователя
FROM ubuntu:latest AS user

RUN useradd -u 10001 go_user

# Финальный образ
FROM scratch AS release

# Copy builded binary
COPY --from=ghcr.io/tarampampam/curl:8.6.0 /bin/curl /bin/curl
COPY --from=build --chown=go_user /app/bin/auth ./auth
COPY --from=user /etc/passwd /etc/passwd

EXPOSE 8080
USER go_user
CMD ["./auth"]