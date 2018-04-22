package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/m0t0k1ch1/metamask-login-sample/interfaces"
	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

const (
	DefaultConfigPath = "config.json"
)

func loadConfig(path string) (*server.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var conf server.Config
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return nil, err
	}
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	return &conf, nil
}

func main() {
	var confPath = flag.String("conf", DefaultConfigPath, "path to your config.json")
	flag.Parse()

	conf, err := loadConfig(*confPath)
	if err != nil {
		panic(err)
	}

	srv := interfaces.NewServer(conf)

	done := make(chan bool, 1)
	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGTERM)
		<-sigterm

		if err := srv.Shutdown(context.Background()); err != nil {
			srv.Logger().Fatal(err)
		}
		close(done)
	}()
	if err := srv.Start(); err != nil {
		srv.Logger().Info(err)
	}
	<-done
}
