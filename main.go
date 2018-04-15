package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	srv := server.New(conf, config.NewContainer())

	done := make(chan bool, 1)

	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGTERM)
		<-sigterm

		if err := srv.Shutdown(context.Background()); err != nil {
			srv.Logger.Fatal(err)
		}
		close(done)
	}()

	if err := srv.Start(); err != nil {
		srv.Logger.Info(err)
	}
	<-done
}
