package main

import (
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/cache"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

func main() {
	conf := config.NewConfig()
	container := &domain.Container{
		NewUserRepository: cache.NewUserRepository,
	}

	srv := server.New(conf, container)

	srv.Logger.Fatal(srv.Start())
}
