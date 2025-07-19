#!make
include .env

db-status:
	@goose -dir=$(GOOSE_MIGRATION_PATH) postgres $(GOOSE_DSN) status

db-up:
	@goose -dir=$(GOOSE_MIGRATION_PATH) postgres $(GOOSE_DSN) up

db-down:
	@goose -dir=$(GOOSE_MIGRATION_PATH) postgres $(GOOSE_DSN) down
