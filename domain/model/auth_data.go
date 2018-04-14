package model

import (
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	AuthDataType = "string"
	AuthDataName = "challenge"
)

type AuthData struct {
	Type  string
	Name  string
	Value string
}

func NewAuthData(value string) *AuthData {
	return &AuthData{
		Type:  AuthDataType,
		Name:  AuthDataName,
		Value: value,
	}
}

func (data *AuthData) signatureHashBytes() []byte {
	return crypto.Keccak256(
		crypto.Keccak256([]byte(data.Type+" "+data.Name)),
		crypto.Keccak256([]byte(data.Value)),
	)
}

func (data *AuthData) RecoverPubkey(sig Signature) (*Pubkey, error) {
	pubkey, err := crypto.SigToPub(data.signatureHashBytes(), sig.Bytes())
	if err != nil {
		return nil, err
	}

	return &Pubkey{pubkey}, nil
}
