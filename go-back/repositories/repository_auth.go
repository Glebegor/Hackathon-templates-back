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
