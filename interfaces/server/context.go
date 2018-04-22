package server

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/application/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/application/user"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/auth/metamask"
)

type Context struct {
	echo.Context
	appCreator *AppCreator
	appCore    *application.Core
}

func (c *Context) NewAuthApplication() auth.Application {
	return c.appCreator.Auth(c.appCore)
}

func (c *Context) NewUserApplication() user.Application {
	return c.appCreator.User(c.appCore)
}

func (c *Context) Claims() *metamask.Claims {
	return c.Get("user").(*jwt.Token).Claims.(*metamask.Claims)
}

func (c *Context) JSONSuccess(result interface{}) error {
	return c.JSON(http.StatusOK, NewSuccessResponse(result))
}

func (c *Context) JSONError(err error) error {
	var result *common.Error
	if commonErr, ok := err.(*common.Error); ok {
		result = commonErr
	} else {
		c.Logger().Error(err)
		result = common.NewUnexpectedError()
	}

	return c.JSON(http.StatusOK, NewErrorResponse(result))
}
