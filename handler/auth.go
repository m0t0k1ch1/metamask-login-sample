package handler

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/application/input"
	"github.com/m0t0k1ch1/metamask-login-sample/handler/response"
)

type AuthController struct {
	app *application.AuthApplication
}

func NewAuthController(app *application.AuthApplication) *AuthController {
	return &AuthController{
		app: app,
	}
}

func (controller *AuthController) ChallengeHandler(c echo.Context) error {
	addressHex := c.FormValue("address")

	in := input.NewAuthChallengeInput(addressHex)

	out, err := controller.app.Challenge(in)
	if err != nil {
		return response.JSONError(c, err)
	}

	return response.JSONSuccess(c, out)
}

func (controller *AuthController) AuthorizeHandler(c echo.Context) error {
	addressHex := c.FormValue("address")
	sigHex := c.FormValue("signature")

	in := input.NewAuthAuthorizeInput(addressHex, sigHex)

	out, err := controller.app.Authorize(in)
	if err != nil {
		return response.JSONError(c, err)
	}

	return response.JSONSuccess(c, out)
}
