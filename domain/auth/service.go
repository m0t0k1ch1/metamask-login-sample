package auth

import (
	"time"

	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
)

type Service interface {
	Authorize(u *user.User, sig common.Signature) (*Token, error)
	Sign(token *Token) (string, error)
}

type service struct {
	secret              string
	tokenExpiryDuration time.Duration
}

func NewService(secret string, ted time.Duration) Service {
	return &service{
		secret:              secret,
		tokenExpiryDuration: ted,
	}
}

func (s *service) Authorize(u *user.User, sig common.Signature) (*Token, error) {
	challenge := NewChallenge(u.ChallengeString())

	address, err := challenge.RecoverAddress(sig)
	if err != nil {
		return nil, err
	}
	if address.Hex() != u.Address.Hex() {
		return nil, common.ErrInvalidSignature
	}

	return NewToken(address, s.tokenExpiryDuration), nil
}

func (s *service) Sign(token *Token) (string, error) {
	return token.SignedString(s.secret)
}
