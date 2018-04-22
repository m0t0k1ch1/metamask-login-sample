package interfaces

import (
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	appAuth "github.com/m0t0k1ch1/metamask-login-sample/application/auth"
	appUser "github.com/m0t0k1ch1/metamask-login-sample/application/user"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/auth/metamask"
	cacheUser "github.com/m0t0k1ch1/metamask-login-sample/infrastructure/cache/user"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server/handler"
)

func NewServer(conf *server.Config) *server.Server {
	srv := server.New(newCore(conf))

	srv.File("/", srv.Config.IndexFilePath)
	srv.Static("/static", srv.Config.StaticDirPath)

	handler.SetUp(srv.Base())

	return srv
}

func newCore(conf *server.Config) *server.Core {
	return server.NewCore(
		conf,
		newAppCreator(conf),
		application.NewCore(
			newContainer(conf),
			conf.App,
		),
	)
}

func newAppCreator(conf *server.Config) *server.AppCreator {
	return &server.AppCreator{
		Auth: appAuth.NewApplication,
		User: appUser.NewApplication,
	}
}

func newContainer(conf *server.Config) *domain.Container {
	return &domain.Container{
		Services: &domain.Services{
			Auth: metamask.NewService(
				conf.App.Auth.Secret,
				conf.App.Auth.TokenExpiryDuration(),
			),
		},
		Repositories: &domain.Repositories{
			User: cacheUser.NewRepository(),
		},
	}
}
