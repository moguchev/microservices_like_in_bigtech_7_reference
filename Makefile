include scratch.mk
include vendor.proto.mk

ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

DB_USERS        = "user=$(USERS_POSTGRES_USER)        password=$(USERS_POSTGRES_PASSWORD)        dbname=$(USERS_POSTGRES_DB)        host=127.0.0.1 port=5434 sslmode=disable"
DB_SOCIAL       = "user=$(SOCIAL_POSTGRES_USER)       password=$(SOCIAL_POSTGRES_PASSWORD)       dbname=$(SOCIAL_POSTGRES_DB)       host=127.0.0.1 port=5436 sslmode=disable"
DB_CHAT         = "user=$(CHAT_POSTGRES_USER)         password=$(CHAT_POSTGRES_PASSWORD)         dbname=$(CHAT_POSTGRES_DB)         host=127.0.0.1 port=5435 sslmode=disable"
DB_NOTIFICATION = "user=$(NOTIFICATION_POSTGRES_USER) password=$(NOTIFICATION_POSTGRES_PASSWORD) dbname=$(NOTIFICATION_POSTGRES_DB) host=127.0.0.1 port=5437 sslmode=disable"

migrate-users:
	@echo "Migrating users service..."
	@goose -dir ./users/migrations postgres $(DB_USERS) up

migrate-social:
	@echo "Migrating social service..."
	@goose -dir ./social/migrations postgres $(DB_SOCIAL) up

migrate-chat:
	@echo "Migrating chat service..."
	@goose -dir ./chat/migrations postgres $(DB_CHAT) up

migrate-notification:
	@echo "Migrating notification service..."
	@goose -dir ./notification/migrations postgres $(DB_NOTIFICATION) up

# migrate - запустить миграции для всех сервисов
migrate: migrate-users migrate-social migrate-chat migrate-notification
	@echo "All migrations applied successfully!"

# docker-compose
.docker-build:
	docker compose build

.docker-up:
	docker compose up -d

# up - поднять локально все
up: .docker-build .docker-up

# down - остановить локально все
down:
	docker compose down

# Объявляем, что текущие команды не являются файлами и
# интсрументируем Makefile не искать изменения в файловой системе
.PHONY: \
	migrate-users \
	migrate-social \
	migrate-chat \
	migrate-notification \
	migrate \
	up \
	.docker-up \
	.docker-build \
	down
