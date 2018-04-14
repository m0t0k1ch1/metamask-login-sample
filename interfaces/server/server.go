package server

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/api/users"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/middleware"
)

type Server struct {
	*echo.Echo
	config    *config.Config
	container *domain.Container
}

func New(conf *config.Config, container *domain.Container) *Server {
	srv := &Server{
		Echo:      echo.New(),
		config:    conf,
		container: container,
	}

	srv.Use(middleware.NewLogger(conf))
	srv.File("/", conf.Server.IndexFilePath)
	srv.Static("/static", conf.Server.StaticDirPath)

	authGroup := srv.Group("/auth")
	authGroup.POST("/challenge", srv.newHandlerFunc(auth.ChallengeHandler))
	authGroup.POST("/authorize", srv.newHandlerFunc(auth.AuthorizeHandler))

	apiGroup := srv.Group("/api")
	apiGroup.Use(middleware.NewAuthenticator(conf))
	apiGroup.GET("/users/:address", srv.newHandlerFunc(users.GetHandler))

	return srv
}

func (srv *Server) newHandlerFunc(h handler.Handler) echo.HandlerFunc {
	return func(ec echo.Context) error {
		return h(&handler.Context{
			Context:   ec,
			Config:    srv.config,
			Container: srv.container,
		})
	}
}

func (srv *Server) Start() error {
	return srv.Echo.Start(":" + srv.config.Server.Port)
}
