package auth

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/spec"
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
	if ok := spec.IsValidAddressHex(in.AddressHex); !ok {
		return domain.ErrInvalidAddressHex
	}
	return nil
}

func (in *ChallengeInput) Address() model.Address {
	return model.NewAddressFromHex(in.AddressHex)
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
	if ok := spec.IsValidAddressHex(in.AddressHex); !ok {
		return domain.ErrInvalidAddressHex
	}
	if ok := spec.IsValidSignatureHex(in.SigHex); !ok {
		return domain.ErrInvalidSignatureHex
	}
	return nil
}

func (in *AuthorizeInput) Address() model.Address {
	return model.NewAddressFromHex(in.AddressHex)
}

func (in *AuthorizeInput) Signature() model.Signature {
	return model.NewSignatureFromHex(in.SigHex)
}
