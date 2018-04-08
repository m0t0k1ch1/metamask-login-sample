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

	e.GET("/challenge", getChallenge)

	e.Logger.Fatal(e.Start(":1323"))
}

// GET /challenge
func getChallenge(c echo.Context) error {
	// TODO
	challenge := "poyon"

	return helpers.JSONResponseSuccess(c, map[string]string{
		"challenge": challenge,
	})
}
