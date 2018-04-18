package handler

import (
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/api"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/auth"
)

func SetUp(g *server.Group) {
	auth.SetUp(g.NewGroup("/auth"))
	api.SetUp(g.NewGroup("/api"))
}
