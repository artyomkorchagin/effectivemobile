include .env

build:
	docker compose --env-file ./.env build

up:
	docker compose --env-file ./.env up

down:
	docker compose --env-file ./.env down

restart: down up

test:
	@go test -v ./...

clean:
	docker compose --env-file ./.env down -v --rmi all

cover:
	@go test -v -coverprofile ./tests/cover.out ./...
	@go tool cover -html ./tests/cover.out -o ./tests/cover.html

clean-volumes:
	docker volume prune -f