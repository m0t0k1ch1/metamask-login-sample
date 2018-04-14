package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewLogger() echo.MiddlewareFunc {
	return middleware.Logger()
}
