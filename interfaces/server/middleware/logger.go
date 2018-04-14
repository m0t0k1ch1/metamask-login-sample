package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/config"
)

func NewLogger(conf *config.Config) echo.MiddlewareFunc {
	return middleware.Logger()
}
