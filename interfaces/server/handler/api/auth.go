package api

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
)

func claims(c echo.Context) *model.AuthClaims {
	return c.Get("user").(*jwt.Token).Claims.(*model.AuthClaims)
}

func VerifyUser(c echo.Context, addressHex string) bool {
	return claims(c).AddressHex == addressHex
}
