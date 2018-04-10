package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/handler"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/storage"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.File("/", "index.html")
	e.Static("/static", "static")

	setUpHandlers(e.Router())

	e.Logger.Fatal(e.Start(":1323"))
}

func setUpHandlers(router *echo.Router) {
	userStorage := storage.NewUserStorage()

	authApp := application.NewAuthApplication(userStorage)
	authController := handler.NewAuthController(authApp)

	router.Add(http.MethodPost, "/challenge", authController.ChallengeHandler)
	router.Add(http.MethodPost, "/authorize", authController.AuthorizeHandler)
}
