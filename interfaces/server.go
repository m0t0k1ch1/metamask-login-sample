package interfaces

import (
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/api"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler/auth"
)

type Server struct {
	*echo.Echo
	Config *server.Config
	Core   *application.Core
}

func NewServer(conf *server.Config) *Server {
	srv := &Server{
		Echo:   echo.New(),
		Config: conf,
		Core: application.NewCore(
			server.NewContainer(conf),
			conf.App,
		),
	}

	srv.Logger.SetLevel(srv.Config.LogLvl())

	srv.File("/", srv.Config.IndexFilePath)
	srv.Static("/static", srv.Config.StaticDirPath)

	authGroup := srv.NewGroup("/auth")
	auth.SetUpHandlers(authGroup)

	apiGroup := srv.NewGroup("/api")
	api.SetUpHandlers(apiGroup)

	return srv
}

func (srv *Server) NewGroup(prefix string) *handler.Group {
	return &handler.Group{
		Group:  srv.Echo.Group(prefix),
		Config: srv.Config,
		Core:   srv.Core,
	}
}

func (srv *Server) Start() error {
	return srv.Echo.Start(srv.Config.Address())
}
