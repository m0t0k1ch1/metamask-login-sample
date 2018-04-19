package common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/library/strutil"
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

func ValidateSignatureHex(sigHex string) error {
	if strutil.HasHexPrefix(sigHex) {
		sigHex = sigHex[2:]
	}

	if len(sigHex) != 2*SignatureLength {
		return ErrInvalidSignatureLength
	}
	if !strutil.IsHex(sigHex) {
		return ErrInvalidSignatureHex
	}

	return nil
}
