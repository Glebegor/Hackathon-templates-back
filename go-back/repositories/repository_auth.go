package repositories

import (
	"project-hackathon/core/common"
	"project-hackathon/core/domain"

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
func (r *repositoryAuth) CheckUserByEmailAndPassword(user *common.UserLogin) (domain.User, error) {
	var u domain.User
	err := r.db.Get(&u, "SELECT id, name, email, password_hash FROM users WHERE email = $1 AND password_hash = $2", user.Email, user.Password)
	return u, err
}
func (r *repositoryAuth) GetUserById(id string) (domain.User, error) {
	var u domain.User
	err := r.db.Get(&u, "SELECT id, name, email, password_hash FROM users WHERE id = $1", id)
	return u, err
}
