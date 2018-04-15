package auth

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
)

const (
	ChallengeType = "string"
	ChallengeName = "challenge"
)

type Challenge struct {
	Type  string
	Name  string
	Value string
}

func NewChallenge(value string) *Challenge {
	return &Challenge{
		Type:  ChallengeType,
		Name:  ChallengeName,
		Value: value,
	}
}

func (challenge *Challenge) signatureHashBytes() []byte {
	return crypto.Keccak256(
		crypto.Keccak256([]byte(challenge.Type+" "+challenge.Name)),
		crypto.Keccak256([]byte(challenge.Value)),
	)
}

func (challenge *Challenge) RecoverAddress(sig common.Signature) (common.Address, error) {
	pubkey, err := crypto.SigToPub(challenge.signatureHashBytes(), sig.Bytes())
	if err != nil {
		return common.Address{}, err
	}

	return common.Address(crypto.PubkeyToAddress(*pubkey)), nil
}
