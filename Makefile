#!make
include .env

PROJECT_NAME=effectivemobile

COMPOSE_FILE=docker-compose.yml

db-status:
	@goose -dir=$(GOOSE_MIGRATION_PATH) postgres $(GOOSE_DSN) status

db-up:
	@goose -dir=$(GOOSE_MIGRATION_PATH) postgres $(GOOSE_DSN) up

db-down:
	@goose -dir=$(GOOSE_MIGRATION_PATH) postgres $(GOOSE_DSN) down

build:
    docker-compose build

up:
    docker-compose up -d

down:
    docker-compose down

restart: down up

migrate:
    docker-compose up -d db
    @sleep 3
    docker-compose up migrate

run:
    go run cmd/main.go

test:
    go test -v ./...

clean:
    docker-compose down -v --rmi all