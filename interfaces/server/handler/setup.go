package handler

import (
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/api"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/auth"
)

func SetUp(cntl *server.Controller) {
	auth.SetUp(cntl.Child("/auth"))
	api.SetUp(cntl.Child("/api"))
}
