package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	appAuth "github.com/m0t0k1ch1/metamask-login-sample/application/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
	dbUser "github.com/m0t0k1ch1/metamask-login-sample/infrastructure/db/user"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/handler/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/handler/users"
)

const (
	DefaultServerPort = "1323"
	DefaultSecret     = "secret"
)

type config struct {
	port   string
	secret string
}

func newConfig() *config {
	port := os.Getenv("MLS_SERVER_PORT")
	if port == "" {
		port = DefaultServerPort
	}

	secret := os.Getenv("MLS_SECRET")
	if secret == "" {
		secret = DefaultSecret
	}

	return &config{
		port:   port,
		secret: secret,
	}
}

func injectDependencies(config *config) {
	appAuth.Secret = func() string {
		return config.secret
	}

	user.NewRepository = dbUser.NewRepository
}

func main() {
	config := newConfig()

	injectDependencies(config)

	e := echo.New()
	e.Use(middleware.Logger())
	e.File("/", "index.html")
	e.Static("/static", "static")

	e.POST("/challenge", auth.ChallengeHandler)
	e.POST("/authorize", auth.AuthorizeHandler)

	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &domain.AuthClaims{},
		SigningKey: []byte(config.secret),
	}))
	apiGroup.GET("/users/:address", users.GetHandler)

	e.Logger.Fatal(e.Start(":" + config.port))
}
