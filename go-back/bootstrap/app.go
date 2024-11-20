package bootstrap

import (
	"github.com/jmoiron/sqlx"
)

type Application struct {
	Env      *Env
	Database *sqlx.DB
}

func App() (*Application, error) {

	app := &Application{}
	env, err := NewEnv()

	if err != nil {
		return app, err
	} else {
		app.Env = env
	}

	db, err := ConnectDatabase(env)

	if err != nil {
		return app, err
	} else {
		app.Database = db
	}

	return app, nil
}
