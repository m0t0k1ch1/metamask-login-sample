package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/handler"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/storage"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.File("/", "index.html")
	e.Static("/static", "static")

	authAPI := handler.NewAuthAPI(storage.NewUserStorage())
	e.POST("/challenge", authAPI.ChallengeHandler)
	e.POST("/login", authAPI.LoginHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
