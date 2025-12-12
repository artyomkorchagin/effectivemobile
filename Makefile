include .env

build:
	docker compose --env-file ./.env build --no-cache

up:
	docker compose --env-file ./.env up -d

down:
	docker compose --env-file ./.env down

restart: down build up

run:
	docker compose --env-file ./.env up

test:
	@go test -v ./...

cover:
	@mkdir -p ./tests
	@go test -v -coverprofile ./tests/cover.out ./...
	@go tool cover -html ./tests/cover.out -o ./tests/cover.html

clean:
	docker compose --env-file ./.env down -v --remove-orphans
	docker compose --env-file ./.env rm -fsv
	docker volume prune -f
	docker image prune -f

clean-volumes:
	docker volume prune -f

.PHONY: build up down restart run test cover clean clean-volumes