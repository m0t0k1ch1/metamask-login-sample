package common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/library/strutil"
)

type Signature [AuthSignatureLength]byte

func (sig Signature) Bytes() []byte {
	return sig[:]
}

func NewSignatureFromHex(sigHex string) Signature {
	sig := Signature{}

	copy(sig[:], common.FromHex(sigHex))

	switch sig[AuthSignatureLength-1] {
	case 27:
		sig[AuthSignatureLength-1] = 0
	case 28:
		sig[AuthSignatureLength-1] = 1
	}

	return sig
}

func ValidateSignatureHex(sigHex string) error {
	if strutil.HasHexPrefix(sigHex) {
		sigHex = sigHex[2:]
	}

	if len(sigHex) != 2*AuthSignatureLength {
		return ErrInvalidSignatureLength
	}
	if !strutil.IsHex(sigHex) {
		return ErrInvalidSignatureHex
	}

	return nil
}
