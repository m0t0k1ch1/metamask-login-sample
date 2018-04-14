package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
)

type Token struct {
	*jwt.Token
}

func NewToken(address common.Address, duration time.Duration) *Token {
	return &Token{
		jwt.NewWithClaims(
			jwt.SigningMethodHS256, NewClaims(address, duration),
		),
	}
}

func (token *Token) SignedString(secret string) (string, error) {
	return token.Token.SignedString([]byte(secret))
}
