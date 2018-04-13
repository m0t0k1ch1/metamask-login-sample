package domain

import (
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	AuthTypedDataType = "string"
	AuthTypedDataName = "challenge"
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

func (data AuthTypedData) signatureHashBytes() []byte {
	return crypto.Keccak256(
		crypto.Keccak256([]byte(data.Type+" "+data.Name)),
		crypto.Keccak256([]byte(data.Value)),
	)
}

func (data AuthTypedData) RecoverPubkey(sig Signature) (*Pubkey, error) {
	pubkey, err := crypto.SigToPub(data.signatureHashBytes(), sig.Bytes())
	if err != nil {
		return nil, err
	}

	return &Pubkey{pubkey}, nil
}
