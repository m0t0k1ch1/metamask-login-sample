package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/handler/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/handler/users"
)

func newAPIAuthenticator(conf *config.Config) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &domain.AuthClaims{},
		SigningKey: []byte(conf.Secret),
	})
}

func SetUp(e *echo.Echo, conf *config.Config) {
	e.POST("/challenge", auth.ChallengeHandler)
	e.POST("/authorize", auth.AuthorizeHandler)

	apiGroup := e.Group("/api")
	apiGroup.Use(newAPIAuthenticator(conf))
	apiGroup.GET("/users/:address", users.GetHandler)
}
