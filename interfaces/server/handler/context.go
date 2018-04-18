package handler

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/auth/metamask"
)

type Context struct {
	echo.Context
	Core *application.Core
}

func NewContext(ec echo.Context, core *application.Core) *Context {
	return &Context{
		Context: ec,
		Core:    core,
	}
}

func (c *Context) Claims() *metamask.Claims {
	return c.Get("user").(*jwt.Token).Claims.(*metamask.Claims)
}
