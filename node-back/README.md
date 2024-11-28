# Nodejs backend template

## Description
This is application based on express.js

## Commands

You have 3 basic types of env (dev, prod, test)

### Run start

run version local ```npm run start ${your_env}```:

```
    "start": "npx ts-node src/index.ts",
```

### Run and start

To build and run write:

```bash
npm run build-start
```

### Start database

To Create database use:

```bash
npm run db-run-dev
```

### Migrates

Creates .env file with connection string + add type of the env ```npm run db-migrate ${your_env}```:

```json
    "db-migrate": "npx ts-node src/bootstrap/setupenv.ts && npx prisma migrate dev --name init"
```

## technologies

- Express.JS
- Prisma ORM
- Node.js
- Typescript
- Swagger
- Docker
