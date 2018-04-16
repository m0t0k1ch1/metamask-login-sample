package server

import (
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/auth/metamask"
	"github.com/m0t0k1ch1/metamask-login-sample/infrastructure/cache/user"
)

var (
	ErrUnknownLogLevel = errors.New("Unknown log level")
)

var (
	logLevels = map[string]log.Lvl{
		"DEBUG": log.DEBUG,
		"INFO":  log.INFO,
		"WARN":  log.WARN,
		"ERROR": log.ERROR,
		"OFF":   log.OFF,
	}
)

type Config struct {
	Port          int                 `json:"port"`
	IndexFilePath string              `json:"index_file_path"`
	StaticDirPath string              `json:"static_dir_path"`
	LogLevel      string              `json:"log_level"`
	App           *application.Config `json:"app"`
}

func (conf *Config) Validate() error {
	if _, ok := logLevels[conf.LogLevel]; !ok {
		return ErrUnknownLogLevel
	}
	return nil
}

func (conf *Config) Address() string {
	return fmt.Sprintf(":%d", conf.Port)
}

func (conf *Config) LogLvl() log.Lvl {
	return logLevels[conf.LogLevel]
}

func (conf *Config) container() *domain.Container {
	return &domain.Container{
		AuthService: metamask.NewService(
			conf.App.Auth.Secret,
			conf.App.Auth.TokenExpiryDuration(),
		),
		UserRepo: user.NewRepository(),
	}
}
