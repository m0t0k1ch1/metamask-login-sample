package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/handler"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/storage"
)

const (
	DefaultServerPort = "1323"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.File("/", "index.html")
	e.Static("/static", "static")

	setUpHandlers(e.Router())

	e.Logger.Fatal(e.Start(":" + getPort()))
}

func getPort() string {
	port := os.Getenv("MLS_SERVER_PORT")
	if port == "" {
		port = DefaultServerPort
	}

	return port
}

func setUpHandlers(router *echo.Router) {
	userStorage := storage.NewUserStorage()

	authApp := application.NewAuthApplication(userStorage)
	authController := handler.NewAuthController(authApp)

	router.Add(http.MethodPost, "/challenge", authController.ChallengeHandler)
	router.Add(http.MethodPost, "/authorize", authController.AuthorizeHandler)
}
