# Go lang backend

## Description

## How to run

### Manual

### Docker



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
