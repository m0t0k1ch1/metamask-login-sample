package di

import (
	"github.com/m0t0k1ch1/metamask-login-sample/application/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
)

type Container struct {
	AuthSecret        func() string
	NewUserRepository func() user.Repository
}

func (container *Container) Inject() {
	auth.Secret = container.AuthSecret
	user.NewRepository = container.NewUserRepository
}
