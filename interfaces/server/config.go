package server

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/cache/user"
)

type Config struct {
	Port          int                 `json:"port"`
	IndexFilePath string              `json:"index_file_path"`
	StaticDirPath string              `json:"static_dir_path"`
	App           *application.Config `json:"app"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func (conf *Config) Address() string {
	return fmt.Sprintf(":%d", conf.Port)
}

func (conf *Config) newContainer() *domain.Container {
	return &domain.Container{
		AuthService: auth.NewService(
			conf.App.Auth.Secret,
			conf.App.Auth.TokenExpiryDuration(),
		),
		UserRepo: user.NewRepository(),
	}
}
