package domain

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	AuthTypedDataType = "string"
	AuthTypedDataName = "challenge"

	AuthClaimsExpiryDuration = 72 * time.Hour
)

type AuthTypedData struct {
	Type  string
	Name  string
	Value string
}

func NewAuthTypedData(value string) *AuthTypedData {
	return &AuthTypedData{
		Type:  AuthTypedDataType,
		Name:  AuthTypedDataName,
		Value: value,
	}
}

func (data *AuthTypedData) signatureHashBytes() []byte {
	return crypto.Keccak256(
		crypto.Keccak256([]byte(data.Type+" "+data.Name)),
		crypto.Keccak256([]byte(data.Value)),
	)
}

func (data *AuthTypedData) RecoverPubkey(sig Signature) (*Pubkey, error) {
	pubkey, err := crypto.SigToPub(data.signatureHashBytes(), sig.Bytes())
	if err != nil {
		return nil, err
	}

	return &Pubkey{pubkey}, nil
}

type AuthClaims struct {
	Address string `json:"address"`
	jwt.StandardClaims
}

func NewAuthClaims(address Address) *AuthClaims {
	now := time.Now()

	return &AuthClaims{
		Address: address.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(AuthClaimsExpiryDuration).Unix(),
			IssuedAt:  now.Unix(),
		},
	}
}
