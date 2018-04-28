package domain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/library/strutil"
)

type Signature [SignatureSize]byte

func NewSignatureFromBytes(sigBytes []byte) Signature {
	sig := Signature{}
	copy(sig[:], sigBytes[:])

	if sig[SignatureSize-1] < SignatureRIRangeBase {
		sig[SignatureSize-1] += SignatureRIRangeBase
	}

	return sig
}

func NewSignatureFromHex(sigHex string) Signature {
	return NewSignatureFromBytes(common.FromHex(sigHex))
}

func (sig Signature) Bytes() []byte {
	return sig[:]
}

func ValidateSignatureHex(sigHex string) error {
	if strutil.HasHexPrefix(sigHex) {
		sigHex = sigHex[2:]
	}

	if len(sigHex) != 2*SignatureSize {
		return ErrInvalidSignatureSize
	}
	if !strutil.IsHex(sigHex) {
		return ErrInvalidSignatureHex
	}

	return nil
}
