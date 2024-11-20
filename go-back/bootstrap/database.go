package bootstrap

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func ConnectDatabase(env *Env) (*sqlx.DB, error) {
	var db *sqlx.DB
	db, err := sqlx.Connect("postgres", "host="+env.DB_HOST+" port="+env.DB_PORT+" user="+env.DB_USER+" password="+env.DB_PASSWORD+" dbname="+env.DB_NAME+" sslmode="+env.DB_SSLMODE)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}