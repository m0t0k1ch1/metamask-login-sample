package auth

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
)

type ChallengeInput struct {
	AddressHex string
}

func NewChallengeInput(addressHex string) *ChallengeInput {
	return &ChallengeInput{
		AddressHex: addressHex,
	}
}

func (in *ChallengeInput) Validate() error {
	if ok := domain.IsValidAddressHex(in.AddressHex); !ok {
		return domain.ErrInvalidAddressFormat
	}
	return nil
}

func (in *ChallengeInput) Address() domain.Address {
	return domain.NewAddressFromHex(in.AddressHex)
}

type AuthorizeInput struct {
	AddressHex string
	SigHex     string
}

func NewAuthorizeInput(addressHex, sigHex string) *AuthorizeInput {
	return &AuthorizeInput{
		AddressHex: addressHex,
		SigHex:     sigHex,
	}
}

func (in *AuthorizeInput) Validate() error {
	if ok := domain.IsValidAddressHex(in.AddressHex); !ok {
		return domain.ErrInvalidAddressFormat
	}
	if ok := domain.IsValidSignatureHex(in.SigHex); !ok {
		return domain.ErrInvalidSignatureFormat
	}
	return nil
}

func (in *AuthorizeInput) Address() domain.Address {
	return domain.NewAddressFromHex(in.AddressHex)
}

func (in *AuthorizeInput) Signature() domain.Signature {
	return domain.NewSignatureFromHex(in.SigHex)
}
