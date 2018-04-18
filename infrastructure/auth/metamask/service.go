package metamask

import (
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/auth"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
	"github.com/m0t0k1ch1/metamask-login-sample/library/strutil"
)

const (
	challengeStringLength = 32
)

type service struct {
	secret              string
	tokenExpiryDuration time.Duration
}

func NewService(secret string, ted time.Duration) auth.Service {
	return &service{
		secret:              secret,
		tokenExpiryDuration: ted,
	}
}

func (s *service) SetUpChallenge(u *user.User) error {
	u.Challenge = strutil.Rand(challengeStringLength)
	return nil
}

func (s *service) VerifyResponse(u *user.User, responseBytes []byte) error {
	pubkey, err := crypto.SigToPub(
		challenge(u.Challenge).signatureHashBytes(),
		responseBytes,
	)
	if err != nil {
		return err
	}

	address := common.Address(crypto.PubkeyToAddress(*pubkey))
	if address.Hex() != u.Address.Hex() {
		return common.ErrInvalidSignature
	}

	return nil
}

func (s *service) IssueToken(u *user.User) ([]byte, error) {
	return newToken(u.Address, s.tokenExpiryDuration).signedBytes(s.secret)
}
