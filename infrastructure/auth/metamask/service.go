package metamask

import (
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/auth"
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

func (s *service) SetUpChallenge(u *domain.User) error {
	u.Challenge = strutil.Rand(challengeStringLength)
	return nil
}

func (s *service) VerifyResponse(u *domain.User, responseBytes []byte) error {
	if responseBytes[domain.SignatureSize-1] >= domain.SignatureRIRangeBase {
		responseBytes[domain.SignatureSize-1] -= domain.SignatureRIRangeBase
	}

	pubkey, err := crypto.SigToPub(
		challenge(u.Challenge).signatureHashBytes(),
		responseBytes,
	)
	if err != nil {
		return err
	}

	address := domain.Address(crypto.PubkeyToAddress(*pubkey))
	if address.Hex() != u.Address.Hex() {
		return domain.ErrInvalidSignature
	}

	return nil
}

func (s *service) IssueToken(u *domain.User) ([]byte, error) {
	return newToken(u.Address, s.tokenExpiryDuration).signedBytes(s.secret)
}
