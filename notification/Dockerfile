# Этап сборки
FROM golang:1.25 AS build

WORKDIR /app

# Копируем исходники сервиса
COPY notification/ ./

# Копируем общие пакеты
COPY lib/ ../lib/

# Копируем go.mod и go.sum сервиса
COPY notification/go.mod notification/go.sum ./
RUN go mod download

# Собираем бинарный файл нашего приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/notification notification/cmd/

# Добавляем пользователя
FROM ubuntu:latest AS user

RUN useradd -u 10001 go_user

# Финальный образ
FROM scratch AS release

# Copy builded binary
COPY --from=ghcr.io/tarampampam/curl:8.6.0 /bin/curl /bin/curl
COPY --from=build --chown=go_user /app/bin/notification ./notification
COPY --from=user /etc/passwd /etc/passwd

EXPOSE 8080
USER go_user
CMD ["./notification"]