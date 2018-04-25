package server

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/auth/metamask"
)

type Context struct {
	echo.Context
	Apps *Apps
}

func (c *Context) Claims() *metamask.Claims {
	u := c.Get("user")
	if u == nil {
		return nil
	}

	token, ok := u.(*jwt.Token)
	if !ok {
		return nil
	}

	return token.Claims.(*metamask.Claims)
}

func (c *Context) JSONSuccess(result interface{}) error {
	return c.JSON(http.StatusOK, NewSuccessResponse(result))
}

func (c *Context) JSONError(err error) error {
	var result *domain.Error
	if dErr, ok := err.(*domain.Error); ok {
		result = dErr
	} else {
		c.Logger().Error(err)
		result = domain.NewUnexpectedError()
	}

	return c.JSON(http.StatusOK, NewErrorResponse(result))
}
