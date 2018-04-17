package application

import "github.com/m0t0k1ch1/metamask-login-sample/domain"

type Core struct {
	*domain.Container
	Config *Config
}

func NewCore(container *domain.Container, conf *Config) *Core {
	return &Core{
		Container: container,
		Config:    conf,
	}
}
