DB_HOST=postgres-hackathon
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=123321
DB_NAME=postgres

# Utils
wait-db-ready:
	@until pg_isready -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER); do \
		echo "Waiting for PostgreSQL to be ready..."; \
		sleep 2; \
	done

migrate-up:
	make wait-db-ready
	migrate -path ./ -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up