package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
)

const (
	ClaimsExpiryDuration = 72 * time.Hour
)

type Claims struct {
	AddressHex string `json:"address"`
	jwt.StandardClaims
}

func NewClaims(address common.Address) *Claims {
	now := time.Now()

	return &Claims{
		AddressHex: address.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(ClaimsExpiryDuration).Unix(),
			IssuedAt:  now.Unix(),
		},
	}
}
