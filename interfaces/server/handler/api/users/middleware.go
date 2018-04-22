package users

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

func NewVerifier() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			c := &server.Context{
				Context: ec,
			}
			if c.Param("address") != c.Claims().AddressHex {
				return echo.ErrNotFound
			}
			return h(ec)
		}
	}
}
