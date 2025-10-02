# Директория, в которой хранятся исполняемые файлы проекта и зависимости, необходимые для сборки.
LOCAL_BIN := $(CURDIR)/bin

# Путь до buf
BUF_BIN := $(LOCAL_BIN)/buf

# Установить зависимости, необходимые для генерации кода
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

# Список сервисов c proto файлами
SERVICES := auth chat social users

# Форматирование proto файлов
.buf-format:
	$(info Running buf format for all services...)
	@for service in $(SERVICES); do \
		$(BUF_BIN) format -w $$service/proto; \
	done

# Генерация .pb файлов
.buf-generate:
	$(info Running buf generate for all services...)
	@for service in $(SERVICES); do \
		PATH="$(LOCAL_BIN):$(PATH)" $(BUF_BIN) generate --template $$service/buf.gen.yaml $$service; \
	done

# Список go модулей
MODULES := lib auth chat gateway social users

# go mod tidy для всех сервисов
.tidy: GOBIN=$(LOCAL_BIN)
.tidy:
	$(info Running go mod tidy for all services...)
	@for module in $(MODULES); do \
  		echo "Running go mod tidy for $$module..."; \
		(cd $$module && go mod tidy); \
	done

# Запустить генерацию кода из proto-файлов
generate: .bin-deps .buf-format .buf-generate .tidy
fast-generate: .buf-format .buf-generate .tidy

# Линтер protobuf файлов
.buf-lint:
	$(info run buf lint...)

	$(LOCAL_BIN)/buf lint

# Линтер
lint: .buf-lint

.PHONY: \
	.bin-deps \
	.buf-generate \
	.buf-format \
	.buf-lint \
	.tidy \
	generate \
	lint