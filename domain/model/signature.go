package model

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	SignatureLength = 65
)

type Signature [SignatureLength]byte

func (sig Signature) Bytes() []byte {
	return sig[:]
}

func NewSignatureFromHex(sigHex string) Signature {
	sig := Signature{}

	copy(sig[:], common.FromHex(sigHex))

	switch sig[SignatureLength-1] {
	case 27:
		sig[SignatureLength-1] = 0
	case 28:
		sig[SignatureLength-1] = 1
	}

	return sig
}
