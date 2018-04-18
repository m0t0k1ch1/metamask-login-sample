package api

import (
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/api/users"
)

func SetUp(g *server.Group) {
	authenticator := NewAuthenticator(g.Config.App.Auth.Secret)
	g.Use(authenticator)

	users.SetUp(g.NewGroup("/users"))
}
