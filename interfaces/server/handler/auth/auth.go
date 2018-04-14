package auth

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/response"
)

func ChallengeHandler(c echo.Context) error {
	addressHex := c.FormValue("address")

	app := auth.NewApplication()

	ctx := c.Request().Context()
	in := auth.NewChallengeInput(addressHex)

	out, err := app.Challenge(ctx, in)
	if err != nil {
		return response.JSONError(c, err)
	}

	return response.JSONSuccess(c, out)
}

func AuthorizeHandler(c echo.Context) error {
	addressHex := c.FormValue("address")
	sigHex := c.FormValue("signature")

	app := auth.NewApplication()

	ctx := c.Request().Context()
	in := auth.NewAuthorizeInput(addressHex, sigHex)

	out, err := app.Authorize(ctx, in)
	if err != nil {
		return response.JSONError(c, err)
	}

	return response.JSONSuccess(c, out)
}
