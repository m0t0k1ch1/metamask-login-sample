package handler

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
)

type Context struct {
	echo.Context
	Config    *config.Config
	Container *domain.Container
}
