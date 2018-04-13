package crypto

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func PubkeyToAddress(pubkey *ecdsa.PublicKey) common.Address {
	return crypto.PubkeyToAddress(*pubkey)
}

func PubkeyToAddressHex(pubkey *ecdsa.PublicKey) string {
	return PubkeyToAddress(pubkey).Hex()
}

func RecoverTypedSignature(hashBytes, sigBytes []byte) (*ecdsa.PublicKey, error) {
	return crypto.SigToPub(hashBytes, sigBytes)
}
