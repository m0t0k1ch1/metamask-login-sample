package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/helpers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.File("/", "index.html")
	e.Static("/static", "static")

	e.POST("/challenge", challengeHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

// POST /challenge
func challengeHandler(c echo.Context) error {
	address := c.FormValue("address")
	// TODO: validate address format

	// TODO: generate & save one-time token
	challenge := address

	return helpers.JSONResponseSuccess(c, map[string]string{
		"challenge": challenge,
	})
}
