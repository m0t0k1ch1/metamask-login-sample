package server

import (
	"github.com/m0t0k1ch1/metamask-login-sample/application/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/application/user"
)

type Apps struct {
	Auth auth.Application
	User user.Application
}
