package config

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/cache/user"
)

func NewContainer() *domain.Container {
	return &domain.Container{
		NewUserRepository: user.NewRepository,
	}
}
