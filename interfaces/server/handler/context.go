package handler

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/auth"
)

type Context struct {
	echo.Context
	Core *application.Core
}

func (c *Context) claims() *auth.Claims {
	return c.Get("user").(*jwt.Token).Claims.(*auth.Claims)
}

func (c *Context) VerifyUser(addressHex string) bool {
	return c.claims().AddressHex == addressHex
}
