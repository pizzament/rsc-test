# Запуск локальный и в докере
.PHONY: dc run test

dc:
	docker-compose up  --remove-orphans --build

run:
	CONFIG_FILE=configs/values_local.yaml go run cmd/server/main.go

test:
	go test -race ./...

# Установка goose
.PHONY: install-goose

LOCAL_BIN:=$(CURDIR)/bin
MIGRATION_NAME ?= create_tables

install-goose:
	$(info Installing goose binary into [$(LOCAL_BIN)]...)
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.24.1

create-migration-file:
	$(LOCAL_BIN)/goose -dir migrations create -s $(MIGRATION_NAME) sql

up-migrations:
	$(LOCAL_BIN)/goose -dir migrations postgres "postgresql://rsc-user:rsc-password@localhost:5432/rsc_db?sslmode=disable" up

down-migrations:
	$(LOCAL_BIN)/goose -dir migrations postgres "postgresql://rsc-user:rsc-password@localhost:5432/rsc_db?sslmode=disable" down