package user

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
	"github.com/m0t0k1ch1/metamask-login-sample/library/strutil"
)

const (
	UserChallengeLength = 32
)

type User struct {
	Name      string
	Address   common.Address
	challenge string
}

func NewUser(name string, address common.Address) *User {
	u := &User{
		Name:    name,
		Address: address,
	}
	u.RegenerateChallengeString()

	return u
}

func (u *User) ChallengeString() string {
	return u.challenge
}

func (u *User) RegenerateChallengeString() {
	u.challenge = strutil.Rand(UserChallengeLength)
}
