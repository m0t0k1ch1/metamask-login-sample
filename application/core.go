package application

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
)

type Core struct {
	Config       *Config
	Services     *Services
	Repositories *Repositories
}

type Services struct {
	Auth auth.Service
}

type Repositories struct {
	User user.Repository
}
