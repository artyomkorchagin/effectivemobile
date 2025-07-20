include .env

build:
	docker compose build

up:
	docker compose up

down:
	docker compose down

restart: down up

run:
	go run cmd/main.go

clean:
	docker compose down -v --rmi all

test:
	docker compose up -d db
	docker compose run --rm test
	down