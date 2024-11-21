package bootstrap

import (
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func ConnectDatabase(env *Env) (*sqlx.DB, error) {
	var db *sqlx.DB
	if os.Getenv("GO_CONTAINER") == "docker" {
		env.DB_PORT = "5432"
		env.DB_HOST = env.DB_NAME
	}
	db, err := sqlx.Connect("postgres", "host="+env.DB_HOST+" port="+env.DB_PORT+" user="+env.DB_USER+" password="+env.DB_PASSWORD+" dbname="+env.DB_NAME+" sslmode="+env.DB_SSLMODE)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}
