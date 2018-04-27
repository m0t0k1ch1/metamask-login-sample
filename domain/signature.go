package domain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/library/strutil"
)

type Signature [SignatureLength]byte

func NewSignatureFromBytes(sigBytes []byte) Signature {
	sig := Signature{}
	copy(sig[:], sigBytes[:])

	return sig
}

func NewSignatureFromHex(sigHex string) Signature {
	return NewSignatureFromBytes(common.FromHex(sigHex))
}

func (sig Signature) Bytes() []byte {
	return sig[:]
}

// RI: Recovery Identifier
func (sig *Signature) SwitchToLowerRIRange() {
	if sig[SignatureLength-1] >= SignatureRIRangeBase {
		sig[SignatureLength-1] -= SignatureRIRangeBase
	}
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
