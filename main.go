package main

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/helpers"
)

var (
	ErrInvalidSignature = helpers.NewResponseError(
		http.StatusInternalServerError,
		"Invalid signature",
	)
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.File("/", "index.html")
	e.Static("/static", "static")

	e.POST("/challenge", challengeHandler)
	e.POST("/login", loginHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

// POST /challenge
func challengeHandler(c echo.Context) error {
	addressHex := c.FormValue("address")
	// TODO: validate address format

	// TODO: generate & save one-time challenge
	challenge := addressHex

	return helpers.JSONResponseSuccess(c, map[string]string{
		"challenge": challenge,
	})
}

// POST /login
func loginHandler(c echo.Context) error {
	addressHex := c.FormValue("address")
	sigHex := c.FormValue("signature")
	// TODO: validate address format
	// TODO: validate signature format

	// TODO: fetch one-time challenge
	data := helpers.NewMyTypedData(addressHex)

	pubkey, err := helpers.RecoverTypedSignature(
		data.SignatureHashBytes(),
		common.FromHex(sigHex),
	)
	if err != nil {
		return err // TODO: convert to ResponseError
	}

	if crypto.PubkeyToAddress(*pubkey).Hex() != addressHex {
		return ErrInvalidSignature
	}

	// TODO: create JWT
	token := "success"

	return helpers.JSONResponseSuccess(c, map[string]string{
		"token": token,
	})
}
