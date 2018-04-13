package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
	dbUser "github.com/m0t0k1ch1/metamask-login-sample/infrastructure/db/user"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/handler/auth"
)

const (
	DefaultServerPort = "1323"
)

func getPort() string {
	port := os.Getenv("MLS_SERVER_PORT")
	if port == "" {
		port = DefaultServerPort
	}

	return port
}

func injectDependencies() {
	user.NewRepository = dbUser.NewRepository
}

func main() {
	injectDependencies()

	e := echo.New()

	e.Use(middleware.Logger())

	e.File("/", "index.html")
	e.Static("/static", "static")

	e.POST("/challenge", auth.ChallengeHandler)
	e.POST("/authorize", auth.AuthorizeHandler)

	e.Logger.Fatal(e.Start(":" + getPort()))
}
