package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	AuthClaimsExpiryDuration = 72 * time.Hour
)

type AuthClaims struct {
	AddressHex string `json:"address"`
	jwt.StandardClaims
}

func NewAuthClaims(address Address) *AuthClaims {
	now := time.Now()

	return &AuthClaims{
		AddressHex: address.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(AuthClaimsExpiryDuration).Unix(),
			IssuedAt:  now.Unix(),
		},
	}
}
