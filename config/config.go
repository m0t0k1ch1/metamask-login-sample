package config

import (
	"encoding/json"
	"os"

	"github.com/m0t0k1ch1/metamask-login-sample/interfaces/server"
)

func NewServerConfig(path string) (*server.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var conf server.Config
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
