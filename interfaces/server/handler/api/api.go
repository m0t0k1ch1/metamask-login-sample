package api

import (
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/api/users"
)

func SetUpHandlers(g *handler.Group) {
	authenticator := newAuthenticator(g.Config.App.Auth.Secret)
	g.Use(authenticator)

	usersGroup := g.NewGroup("/users")
	users.SetUpHandlers(usersGroup)
}
