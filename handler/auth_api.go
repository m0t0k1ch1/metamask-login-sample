package handler

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	repo "github.com/m0t0k1ch1/metamask-login-sample/domain/repository"
)

type AuthAPI struct {
	userRepo repo.UserRepository
}

func NewAuthAPI(repo repo.UserRepository) *AuthAPI {
	return &AuthAPI{
		userRepo: repo,
	}
}

func (api *AuthAPI) ChallengeHandler(c echo.Context) error {
	addressHex := c.FormValue("address")

	app := application.NewAuthApplication(api.userRepo)

	challenge, err := app.Challenge(addressHex)
	if err != nil {
		return err
	}

	return jsonResponseSuccess(c, map[string]string{
		"challenge": challenge,
	})
}

func (api *AuthAPI) LoginHandler(c echo.Context) error {
	addressHex := c.FormValue("address")
	sigHex := c.FormValue("signature")

	app := application.NewAuthApplication(api.userRepo)

	token, err := app.Login(addressHex, sigHex)
	if err != nil {
		return err
	}

	return jsonResponseSuccess(c, map[string]string{
		"token": token,
	})
}
