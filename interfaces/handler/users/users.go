package users

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application/user"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/response"
)

func GetHandler(c echo.Context) error {
	addressHex := c.Param("address")

	if ok := auth.VerifyUser(c, addressHex); !ok {
		return echo.ErrNotFound
	}

	app := user.NewApplication()

	ctx := c.Request().Context()
	in := user.NewGetUserInput(addressHex)

	out, err := app.GetUser(ctx, in)
	if err != nil {
		return response.JSONError(c, err)
	}

	return response.JSONSuccess(c, out)
}
