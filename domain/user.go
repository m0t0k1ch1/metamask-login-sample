package domain

import (
	"github.com/m0t0k1ch1/metamask-login-sample/library/strutil"
)

const (
	TokenLength = 32
)

type User struct {
	Name    string
	Address Address
	token   string
}

func NewUser(address Address) *User {
	return &User{
		Address: address,
	}
}

func (u *User) Token() string {
	return u.token
}

func (u *User) AddressHex() string {
	return u.Address.Hex()
}

func (u *User) UpdateToken() {
	u.token = strutil.Rand(TokenLength)
}
