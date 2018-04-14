package main

import (
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/cache/user"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

func main() {
	conf := config.NewServerConfig()
	container := &domain.Container{
		NewUserRepository: user.NewRepository,
	}

	srv := server.New(conf, container)

	srv.Logger.Fatal(srv.Start())
}
