package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/auth/metamask"
)

func NewAuthenticator(secret string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &metamask.Claims{},
		SigningKey: []byte(secret),
	})
}
