package server

import (
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/application/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/application/user"
)

type Core struct {
	Config     *Config
	appCreator *AppCreator
	appCore    *application.Core
}

type AppCreator struct {
	Auth auth.Creator
	User user.Creator
}

func NewCore(conf *Config, appCreator *AppCreator, appCore *application.Core) *Core {
	return &Core{
		Config:     conf,
		appCreator: appCreator,
		appCore:    appCore,
	}
}
