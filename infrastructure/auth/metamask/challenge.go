package metamask

import (
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	challengeType = "string"
	challengeName = "challenge"
)

type challenge struct {
	t string
	n string
	v string
}

func newChallenge(value string) *challenge {
	return &challenge{
		t: challengeType,
		n: challengeName,
		v: value,
	}
}

func (challenge *challenge) signatureHashBytes() []byte {
	return crypto.Keccak256(
		crypto.Keccak256([]byte(challenge.t+" "+challenge.n)),
		crypto.Keccak256([]byte(challenge.v)),
	)
}
