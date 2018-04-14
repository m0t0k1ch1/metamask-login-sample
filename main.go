package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	appAuth "github.com/m0t0k1ch1/metamask-login-sample/application/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
	dbUser "github.com/m0t0k1ch1/metamask-login-sample/infrastructure/db/user"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/handler/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/handler/users"
)

func injectDependencies(conf *config.Config) {
	appAuth.Secret = func() string {
		return conf.Secret
	}

	user.NewRepository = dbUser.NewRepository
}

func main() {
	conf := config.NewConfig()

	injectDependencies(conf)

	e := echo.New()
	e.Use(middleware.Logger())
	e.File("/", conf.Server.IndexFilePath)
	e.Static("/static", conf.Server.StaticDirPath)

	e.POST("/challenge", auth.ChallengeHandler)
	e.POST("/authorize", auth.AuthorizeHandler)

	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &domain.AuthClaims{},
		SigningKey: []byte(conf.Secret),
	}))
	apiGroup.GET("/users/:address", users.GetHandler)

	e.Logger.Fatal(e.Start(":" + conf.Server.Port))
}
