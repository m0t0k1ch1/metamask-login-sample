package auth

import "github.com/m0t0k1ch1/metamask-login-sample/domain"

type Service interface {
	SetUpChallenge(u *domain.User) error
	VerifyResponse(u *domain.User, responseBytes []byte) error
	IssueToken(u *domain.User) ([]byte, error)
}
