include .env

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

restart: down up

run:
	go run cmd/main.go

clean:
	docker compose down -v --rmi all