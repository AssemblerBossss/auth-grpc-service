STORAGE_PATH ?= ./storage/sso.db
MIGRATIONS_PATH ?= ./migrations
MIGRATIONS_TABLE ?= migrations

.PHONY: migrate
migrate:
	mkdir -p "$$(dirname -- "$(STORAGE_PATH)")"
	go run ./cmd/migrator/ --storage-path="$(STORAGE_PATH)" --migrations-path="$(MIGRATIONS_PATH)" --migrations-table="$(MIGRATIONS_TABLE)"
