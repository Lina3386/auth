include .env

LOCAL_BIN:= $(CURDIR)/bin

PROTO_DIR       := proto
GENERATED_DIR   := pkg/user
PB_DIR          := $(GENERATED_DIR)
GRPC_DIR        := $(GENERATED_DIR)
CMD_DIR         := cmd

LOCAL_MIGRATION_DIR := $(MIGRATION_DIR)
LOCAL_MIGRATION_DSN := "host=$(DB_HOST) port=$(DB_PORT) dbname=$(DB_NAME) user=$(DB_USER) password=$(DB_PASSWORD) sslmode=disable"
# Прото-файлы
PROTO_FILES     := $(wildcard $(PROTO_DIR)/*.proto)

# Инструменты
PROTOC          := protoc
PROTOC_GEN_GO   := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc

GOBIN           := $(shell go env GOPATH)/bin

.PHONY: all deps generate clean build server client
.PHONY: generate-chat-api local-migration-status create-migration local-migration-up local-migration-down

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

install-deps:
	@mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

install-golangci-lint:
	@mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

dirs:
	@mkdir -p pkg/user
generate:
	make generate-user

generate-user:
	@echo "Generating user API code..."
	@mkdir -p pkg/user
	$(PROTOC) --proto_path=$(PROTO_DIR) \
		--go_out=pkg/user --go_opt=paths=source_relative \
		--go-grpc_out=pkg/user --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/user.proto
	@echo "Generated user in pkg/user"

local migration-status:
	$(LOCAL_BIN)/goose.exe -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) status -v

create-migration:
	$(LOCAL_BIN)/goose.exe -dir $(LOCAL_MIGRATION_DIR) create users sql

local-migration-up:
	$(LOCAL_BIN)/goose.exe -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) up -v

local-migration-down:
	$(LOCAL_BIN)/goose.exe -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) down -v