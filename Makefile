include scratch.mk
include vendor.proto.mk

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
	up \
	.docker-up \
	.docker-build \
	down