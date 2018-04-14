package model

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type Pubkey struct {
	*ecdsa.PublicKey
}

func (pubkey *Pubkey) Address() Address {
	return Address(crypto.PubkeyToAddress(*pubkey.PublicKey))
}
