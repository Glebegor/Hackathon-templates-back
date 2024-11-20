# Go lang backend

## Description

## Requirements

- makefile
- docker
- go
- yq

add logs.txt to path ```./logs/logs.txt```

## How to run

### Manual

This commands can provide to you manual run of the application.

```bash
make test-db-run

make migrate-up

go run cmd/main.go
```

### Docker


## Makefile

build:
run:
build-run:
test-db-run:
test-db-rm:
migrate-up:
migrate-down:
open-logs:
docker-containers-remove:

## Database

Using Basic and needble presets for database structure

### Users

basic table of users:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    password_hash VARCHAR(1000) NOT NULL,
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
Environment: "dev" # dev, release, test

SERVER:
  PORT: 3000
  HOST: localhost

DB:
  HOST: localhost
  PORT: 5432
  NAME: go-back
  USER: postgres
  SSLMODE: disable
```

On config dependence WHOLE application, soo, be carefull.