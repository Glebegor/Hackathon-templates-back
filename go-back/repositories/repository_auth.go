package repositories

import (
	"project-hackathon/core/common"

	"github.com/jmoiron/sqlx"
)

type repositoryAuth struct {
	db *sqlx.DB
}

func NewRepositoryAuth(db *sqlx.DB) common.RepositoryAuth {
	return &repositoryAuth{db}
}

func (r *repositoryAuth) Register(user *common.UserRegister) error {
	_, err := r.db.Exec("INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	return err
}
