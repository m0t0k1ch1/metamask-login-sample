package domain

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
)

type Container struct {
	AuthService auth.Service
	UserRepo    user.Repository
}
