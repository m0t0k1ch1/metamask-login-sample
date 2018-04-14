package main

import (
	"flag"
	"log"

	"github.com/m0t0k1ch1/metamask-login-sample/config"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

const (
	DefaultConfigPath = "config.json"
)

func main() {
	var confPath = flag.String("conf", DefaultConfigPath, "path to your config.json")
	flag.Parse()

	conf, err := config.NewServerConfig(*confPath)
	if err != nil {
		log.Fatal(err)
	}

	container := config.NewContainer()

	srv := server.New(conf, container)
	srv.Logger.Fatal(srv.Start())
}
