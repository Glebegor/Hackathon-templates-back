# Environment variables
PROJECT_PATH = $(PWD)

# Config variables
PROJECT_NAME = $(shell yq '.PROJECT.NAME' config.yml)
PROJECT_DOCKER_NETWORK = $(shell yq '.PROJECT.DOCKER_NETWORK' config.yml)

SERVER_PORT := $(shell yq '.SERVER.PORT' config.yml)
SERVER_HOST := $(shell yq '.SERVER.HOST' config.yml)
SERVER_DOCKER_PORT := $(shell yq '.SERVER.DOCKER_PORT' config.yml)

DB_PORT := $(shell yq '.DB.PORT' config.yml)
DB_HOST := $(shell yq '.DB.HOST' config.yml)
DB_USER := $(shell yq '.DB.USER' config.yml)
DB_SSLMODE := $(shell yq '.DB.SSLMODE' config.yml)
DB_NAME := $(shell yq '.DB.NAME' config.yml)
DB_PASSWORD := $(shell grep 'SERVER_SECRET' .env | sed 's/SERVER_SECRET=//')

# Utils
wait-db-ready:
	@until pg_isready -h $(DB_HOST) -p $(DB_PORT) -U $(DB_NAME); do \
		echo "Waiting for PostgreSQL to be ready..."; \
		sleep 2; \
	done

# Database
db-run:
	docker run --name=$(DB_NAME) --network $(PROJECT_DOCKER_NETWORK) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d --rm -p $(DB_PORT):5432 postgres:latest
db-rm:
	docker stop $(DB_NAME)

# Migrate
migrate-up:
	migrate -path ./schemas -database "postgres://$(DB_NAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_USER)?sslmode=$(DB_SSLMODE)" up
migrate-down:
	migrate -path ./schemas -database "postgres://$(DB_NAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_USER)?sslmode=$(DB_SSLMODE)" down

# Run docker
build:
	docker image rm -f $(PROJECT_NAME) || true
	docker build -t $(PROJECT_NAME) .
run-d:
	docker network inspect $(PROJECT_DOCKER_NETWORK) >/dev/null 2>&1 || docker network create $(PROJECT_DOCKER_NETWORK)
	make db-run || true
	
	make wait-db-ready
	make migrate-up

	docker rm -f $(PROJECT_NAME) || true	
	docker run --name=$(PROJECT_NAME) -e GO_CONTAINER="docker" --network $(PROJECT_DOCKER_NETWORK) -d --rm -p $(SERVER_DOCKER_PORT):$(SERVER_PORT) $(PROJECT_NAME)
run:
	docker network inspect $(PROJECT_DOCKER_NETWORK) >/dev/null 2>&1 || docker network create $(PROJECT_DOCKER_NETWORK)
	make db-run || true
	
	make wait-db-ready
	make migrate-up

	docker rm -f $(PROJECT_NAME) || true	
	docker run --name=$(PROJECT_NAME) -e GO_CONTAINER="docker" --network $(PROJECT_DOCKER_NETWORK) --rm -p $(SERVER_DOCKER_PORT):$(SERVER_PORT) $(PROJECT_NAME)
build-run:
	make build
	make run
build-run-d:
	make build
	make run-d
stop:
	docker rm -f $(PROJECT_NAME) || true
	make db-rm || true

# Run cluster
wait-db-ready-cluster:
	@until pg_isready -h postgres-hackathon -p 5432 -U postgres; do \
		echo "Waiting for PostgreSQL to be ready..."; \
		sleep 2; \
	done

docker-cluster-run:
	docker build -t database-migration ./schemas 
	docker build -t $(PROJECT_NAME) .
	docker compose up 
docker-cluster-stop:
	docker compose down 
inside-cluster-run:
	make wait-db-ready-cluster
	/main

# help commands
open-logs:
	cat ./logs/logs.txt
docs-update:
	swag init -g cmd/main.go

