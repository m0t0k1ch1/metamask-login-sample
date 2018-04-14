package model

import jwt "github.com/dgrijalva/jwt-go"

type AuthToken struct {
	*jwt.Token
}

func NewAuthToken(address Address) *AuthToken {
	return &AuthToken{
		jwt.NewWithClaims(
			jwt.SigningMethodHS256, NewAuthClaims(address),
		),
	}
}

func (token *AuthToken) SignedString(secret string) (string, error) {
	return token.Token.SignedString([]byte(secret))
}
