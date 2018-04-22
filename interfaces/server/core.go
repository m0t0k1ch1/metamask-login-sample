package server

import (
	"github.com/m0t0k1ch1/metamask-login-sample/application"
)

type Core struct {
	Config     *Config
	AppCreator *AppCreator
	AppCore    *application.Core
}
