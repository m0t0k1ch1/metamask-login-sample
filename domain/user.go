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

func (user *User) Token() string {
	return user.token
}

func (user *User) UpdateToken() {
	user.token = strutil.Rand(TokenLength)
}

func (user *User) AuthTypedData() *AuthTypedData {
	return NewAuthTypedData(user.Token())
}
