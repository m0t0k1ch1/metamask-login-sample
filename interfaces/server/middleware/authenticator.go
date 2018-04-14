package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
)

func NewAuthenticator(conf *config.Config) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &model.AuthClaims{},
		SigningKey: []byte(conf.App.Secret),
	})
}
