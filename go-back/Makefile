# Environment variables
PROJECT_NAME = $(shell basename $(PWD))
PROJECT_PATH = $(PWD)

# Config variables
SERVER_PORT := $(shell yq '.SERVER.PORT' config.yml)
SERVER_HOST := $(shell yq '.SERVER.HOST' config.yml)
DB_PORT := $(shell yq '.DB.PORT' config.yml)
DB_HOST := $(shell yq '.DB.HOST' config.yml)
DB_USER := $(shell yq '.DB.USER' config.yml)
DB_SSLMODE := $(shell yq '.DB.SSLMODE' config.yml)
DB_NAME := $(shell yq '.DB.NAME' config.yml)
DB_PASSWORD := $(shell grep 'SERVER_SECRET' .env | sed 's/SERVER_SECRET=//')




build:
run:
build-run:

test-db-run:
	docker run --name=$(DB_NAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d --rm -p $(DB_PORT):5432 postgres
test-db-rm:
	docker stop $(DB_NAME)

migrate-up:
	migrate -path ./schemas -database "postgres://$(DB_NAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_USER)?sslmode=$(DB_SSLMODE)" up
migrate-down:
	migrate -path ./schemas -database "postgres://$(DB_NAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_USER)?sslmode=$(DB_SSLMODE)" down

open-logs:
	cat ./logs/logs.txt

docker-containers-remove:
