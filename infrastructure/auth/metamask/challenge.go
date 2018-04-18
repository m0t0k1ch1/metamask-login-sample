package metamask

import (
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	challengeType = "string"
	challengeName = "challenge"
)

type challenge string

func (chal challenge) signatureHashBytes() []byte {
	return crypto.Keccak256(
		crypto.Keccak256([]byte(challengeType+" "+challengeName)),
		crypto.Keccak256([]byte(chal)),
	)
}
