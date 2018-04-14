package model

import (
	"github.com/m0t0k1ch1/metamask-login-sample/library/strutil"
)

const (
	UserChallengeLength = 32
)

type User struct {
	Name      string
	Address   Address
	challenge string
}

func NewUser(address Address) *User {
	return &User{
		Address: address,
	}
}

func (user *User) Challenge() string {
	return user.challenge
}

func (user *User) UpdateChallenge() {
	user.challenge = strutil.Rand(UserChallengeLength)
}

func (user *User) AuthData() *AuthData {
	return NewAuthData(user.Challenge())
}
