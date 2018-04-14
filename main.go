package main

import (
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/config/di"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/cache/user"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

func main() {
	conf := config.NewConfig()

	container := &di.Container{
		AuthSecret:        func() string { return conf.Secret },
		NewUserRepository: user.NewRepository,
	}
	container.Inject()

	srv := server.New(conf)
	srv.Logger.Fatal(srv.Start())
}
