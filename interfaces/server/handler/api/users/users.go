package users

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application/user"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/response"
)

func GetHandler(c *handler.Context) error {
	addressHex := c.Param("address")

	if c.Claims().AddressHex != addressHex {
		return echo.ErrNotFound
	}

	app := user.NewApplication(c.Core)

	ctx := c.Request().Context()
	in := user.NewGetUserInput(addressHex)

	out, err := app.GetUser(ctx, in)
	if err != nil {
		return response.JSONError(c, err)
	}

	return response.JSONSuccess(c, out)
}
