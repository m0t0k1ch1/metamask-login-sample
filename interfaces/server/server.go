package server

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/config/di"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/api/users"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/middleware"
)

type Server struct {
	*echo.Echo
	config *config.Config
	di     *di.Container
}

func New(conf *config.Config) *Server {
	e := echo.New()
	e.Use(middleware.NewLogger(conf))
	e.File("/", conf.Server.IndexFilePath)
	e.Static("/static", conf.Server.StaticDirPath)

	authGroup := e.Group("/auth")
	authGroup.POST("/challenge", auth.ChallengeHandler)
	authGroup.POST("/authorize", auth.AuthorizeHandler)

	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.NewAuthenticator(conf))
	apiGroup.GET("/users/:address", users.GetHandler)

	return &Server{
		Echo:   e,
		config: conf,
	}
}

func (srv *Server) Start() error {
	return srv.Echo.Start(":" + srv.config.Server.Port)
}
