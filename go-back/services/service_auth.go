package services

import (
	"project-hackathon/bootstrap"
	"project-hackathon/core/common"
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
