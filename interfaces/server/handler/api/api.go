package api

import (
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/api/users"
)

func SetUp(cntl *server.Controller) {
	authenticator := NewAuthenticator(cntl.Config.App.Auth.Secret)
	cntl.Use(authenticator)

	users.SetUp(cntl.Child("/users"))
}
