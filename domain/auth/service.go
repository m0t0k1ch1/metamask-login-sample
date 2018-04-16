package auth

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
)

type Service interface {
	SetUpChallenge(u *user.User) error
	VerifyResponse(u *user.User, responseBytes []byte) error
	IssueToken(u *user.User) ([]byte, error)
}
