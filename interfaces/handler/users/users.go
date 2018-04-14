package users

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/response"
)

func GetHandler(c echo.Context) error {
	addressHex := c.Param("address")

	claims := auth.Claims(c)
	if claims.AddressHex != addressHex {
		return echo.ErrNotFound
	}

	// TODO: fetch user
	user := claims

	return response.JSONSuccess(c, user)
}
