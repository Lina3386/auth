PROTO_DIR       := proto
GENERATED_DIR   := pkg/user
PB_DIR          := $(GENERATED_DIR)
GRPC_DIR        := $(GENERATED_DIR)
CMD_DIR         := cmd

# Прото-файлы
PROTO_FILES     := $(wildcard $(PROTO_DIR)/*.proto)

# Инструменты
PROTOC          := protoc
PROTOC_GEN_GO   := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc

# Пути для go install (при необходимости)
GOBIN           := $(shell go env GOPATH)/bin

.PHONY: all deps generate clean build server client

all: deps generate build

# Установить плагины для генерации
deps:
	@echo "Устанавливаем protoc-gen-go и protoc-gen-go-grpc..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Создать директории для сгенерированного кода
dirs:
	@mkdir -p $(PB_DIR)
	@mkdir -p $(GRPC_DIR)

# Генерация Go-кода из .proto
generate: dirs
	@echo "Генерируем Go-код из .proto..."
	$(PROTOC) \
	  --proto_path=$(PROTO_DIR) \
	  --go_out=$(PB_DIR) --go_opt=paths=source_relative \
	  --go-grpc_out=$(GRPC_DIR) --go-grpc_opt=paths=source_relative \
	  $(PROTO_FILES)
	@echo "Сгенерировано в $(PB_DIR) и $(GRPC_DIR)"