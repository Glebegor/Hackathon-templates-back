# Get env variables and variables from config.yml
include .env
include config.yml

# Variables
PROJECT_NAME = $(shell basename $(PWD))
PROJECT_PATH = $(PWD)




build:
run:
build-run:

test-db-run:
	docker run --name="test-db-hackathon-project" -e POSTGRES_PASSWORD="123321" -d --rm -p 5436:5432 postgres
test-db-rm:

migrate-up:
migrate-down:

docker-containers-remove:
