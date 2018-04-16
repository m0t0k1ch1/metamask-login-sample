package user

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
)

type User struct {
	Name      string
	Address   common.Address
	Challenge string
}

func NewUser(name string, address common.Address) *User {
	return &User{
		Name:    name,
		Address: address,
	}
}
