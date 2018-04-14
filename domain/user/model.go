package user

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
	"github.com/m0t0k1ch1/metamask-login-sample/library/strutil"
)

const (
	UserChallengeLength = 32
)

type User struct {
	Name      string
	Address   common.Address
	challenge *auth.Challenge
}

func NewUser(address common.Address) *User {
	return &User{
		Address: address,
	}
}

func (user *User) Challenge() *auth.Challenge {
	return user.challenge
}

func (user *User) UpdateChallenge() {
	user.challenge = auth.NewChallenge(strutil.Rand(UserChallengeLength))
}
