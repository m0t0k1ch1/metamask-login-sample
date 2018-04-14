package application

import "github.com/m0t0k1ch1/metamask-login-sample/domain"

type Base struct {
	core *Core
}

func NewBase(core *Core) *Base {
	return &Base{
		core: core,
	}
}

func (base *Base) Config() *Config {
	return base.core.config
}

func (base *Base) Container() *domain.Container {
	return base.core.container
}
