# Go lang backend

## Description

Golang template of backend app for hackathouns.

## Requirements

- makefile
- docker
- go
- yq

add logs.txt to path ```./logs/logs.txt```

## Changes before

1. Change config.yml
2. Change .env
3. Check makefile and containers
4. Check docker-compose
5. Check database

## How to run

### Manual

This commands can provide to you manual run of the application.

```bash
make test-db-run

make migrate-up

go run cmd/main.go
```

### Docker

### Cluster

## Makefile

- build: docker build
- run: docker run
- build-run: docker build and run
- test-db-run: runs creation of db
- test-db-rm: runs deletion of db
- migrate-up: creates migrations
- migrate-down: deletes migrations
- open-logs: showing logs
- docker-cluster-run: runs docker cluster
- docker-cluster-stop: stops docker cluster

## Database

Using Basic and needble presets for database structure.

### Users

basic table of users:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(1000) NOT NULL
);
```

## Config

Env

```dotenv
SERVER_SECRET="123321"
DB_PASSWORD="123321"
```

config.yml

```yml
Environment: "dev"

PROJECT:
  NAME: "go-back-template-hackathon"
  VERSION: "1.0.0"
  DESCRIPTION: "Go back-end project"
  DOCKER_NETWORK: "go-back-temp-hack-net"

SERVER:
  PORT: "3000"
  HOST: "0.0.0.0"
  DOCKER_PORT: "8070"

DB:
  HOST: "localhost"
  PORT: "5436"
  NAME: "postgres"
  USER: "postgres"
  SSLMODE: "disable"
```

On config dependence WHOLE application, soo, be carefull.

## Routers
Init routes to work with project.

### Auth

#### Login
Create user.

#### Register
Register user.

### Checking

#### Ping
Send request to check api.

#### Ping protected
Send request with header Authorization and "Bearer qwlej..." value to check protected routes.