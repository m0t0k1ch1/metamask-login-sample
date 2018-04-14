package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/m0t0k1ch1/metamask-login-sample/application/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
	dbUser "github.com/m0t0k1ch1/metamask-login-sample/infrastructure/db/user"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/handler"
)

func injectDependencies(conf *config.Config) {
	auth.Secret = func() string { return conf.Secret }
	user.NewRepository = dbUser.NewRepository
}

func main() {
	conf := config.NewConfig()

	injectDependencies(conf)

	e := echo.New()
	e.Use(middleware.Logger())
	e.File("/", conf.Server.IndexFilePath)
	e.Static("/static", conf.Server.StaticDirPath)

	handler.SetUp(e, conf)

	e.Logger.Fatal(e.Start(":" + conf.Server.Port))
}
