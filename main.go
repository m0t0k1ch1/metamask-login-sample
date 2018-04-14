package main

import (
	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

func main() {
	srv := server.New(
		config.NewServerConfig(),
		config.NewContainer(),
	)
	srv.Logger.Fatal(srv.Start())
}
