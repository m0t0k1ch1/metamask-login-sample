package users

import (
	"github.com/m0t0k1ch1/metamask-login-sample/application/user"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

func SetUp(cntl *server.Controller) {
	verifier := NewVerifier()
	cntl.GET("/:address", GetHandler, verifier)
	cntl.PUT("/:address", UpdateHandler, verifier)
	cntl.DELETE("/:address", DeleteHandler, verifier)
}

func GetHandler(c *server.Context) error {
	addressHex := c.Param("address")

	app := user.NewApplication(c.Core)

	ctx := c.Request().Context()
	in := user.NewGetUserInput(addressHex)

	out, err := app.GetUser(ctx, in)
	if err != nil {
		return c.JSONError(err)
	}

	return c.JSONSuccess(out)
}

func UpdateHandler(c *server.Context) error {
	addressHex := c.Param("address")
	name := c.FormValue("name")

	app := user.NewApplication(c.Core)

	ctx := c.Request().Context()
	in := user.NewUpdateUserInput(addressHex, name)

	out, err := app.UpdateUser(ctx, in)
	if err != nil {
		return c.JSONError(err)
	}

	return c.JSONSuccess(out)
}

func DeleteHandler(c *server.Context) error {
	addressHex := c.Param("address")

	app := user.NewApplication(c.Core)

	ctx := c.Request().Context()
	in := user.NewDeleteUserInput(addressHex)

	out, err := app.DeleteUser(ctx, in)
	if err != nil {
		return c.JSONError(err)
	}

	return c.JSONSuccess(out)
}
