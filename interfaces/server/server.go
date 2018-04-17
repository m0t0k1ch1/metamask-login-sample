package server

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/api/users"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/middleware"
)

type Server struct {
	*echo.Echo
	config    *Config
	container *domain.Container
}

func New(conf *Config) *Server {
	srv := &Server{
		Echo:      echo.New(),
		config:    conf,
		container: newContainer(conf),
	}

	srv.Logger.SetLevel(srv.config.LogLvl())

	srv.File("/", srv.config.IndexFilePath)
	srv.Static("/static", srv.config.StaticDirPath)

	authGroup := srv.Group("/auth")
	authGroup.POST("/challenge", auth.ChallengeHandler)
	authGroup.POST("/authorize", auth.AuthorizeHandler)

	apiGroup := srv.Group("/api")
	apiGroup.Use(middleware.NewAuthenticator(srv.config.App.Auth.Secret))
	apiGroup.GET("/users/:address", users.GetHandler)

	return srv
}

func (srv *Server) Group(prefix string) *Group {
	return &Group{
		group:  srv.Echo.Group(prefix),
		server: srv,
	}
}

func (srv *Server) Start() error {
	return srv.Echo.Start(srv.config.Address())
}
