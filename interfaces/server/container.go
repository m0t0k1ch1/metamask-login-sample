package server

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/auth/metamask"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/cache/user"
)

func newContainer(conf *Config) *domain.Container {
	return &domain.Container{
		Services: &domain.Services{
			Auth: metamask.NewService(
				conf.App.Auth.Secret,
				conf.App.Auth.TokenExpiryDuration(),
			),
		},
		Repositories: &domain.Repositories{
			User: user.NewRepository(),
		},
	}
}
