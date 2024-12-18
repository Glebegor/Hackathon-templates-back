package services

import (
	"project-hackathon/bootstrap"
	"project-hackathon/core/common"
	"project-hackathon/core/domain"
	"project-hackathon/utils"
	"time"
)

type serviceAuth struct {
	repo    common.RepositoryAuth
	env     *bootstrap.Env
	timeout time.Duration
}

func NewServiceAuth(repo common.RepositoryAuth, env *bootstrap.Env, timeout time.Duration) common.ServiceAuth {
	return &serviceAuth{repo, env, timeout}
}

func (s *serviceAuth) Register(user *common.UserRegister) error {
	user.Password = utils.HashPassword(user.Password, s.env.SERVER_SECRET)
	return s.repo.Register(user)
}

func (s *serviceAuth) CheckUserByEmailAndPassword(user *common.UserLogin) (domain.User, error) {
	return s.repo.CheckUserByEmailAndPassword(user)
}

func (s *serviceAuth) GetUserById(id string) (domain.User, error) {
	return s.repo.GetUserById(id)
}
