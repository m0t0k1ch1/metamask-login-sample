package common

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type PublicKey struct {
	*ecdsa.PublicKey
}

func (pubkey *PublicKey) Address() Address {
	return Address(crypto.PubkeyToAddress(*pubkey.PublicKey))
}
