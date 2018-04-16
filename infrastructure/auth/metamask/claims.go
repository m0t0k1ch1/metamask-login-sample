package metamask

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
)

type Claims struct {
	AddressHex string `json:"address"`
	jwt.StandardClaims
}

func newClaims(address common.Address, d time.Duration) *Claims {
	now := time.Now()

	return &Claims{
		AddressHex: address.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(d).Unix(),
			IssuedAt:  now.Unix(),
		},
	}
}
