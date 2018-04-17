package auth

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
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
	if err := common.ValidateAddressHex(in.AddressHex); err != nil {
		return err
	}
	return nil
}

func (in *ChallengeInput) Address() common.Address {
	return common.NewAddressFromHex(in.AddressHex)
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
	if err := common.ValidateAddressHex(in.AddressHex); err != nil {
		return err
	}
	if err := common.ValidateSignatureHex(in.SigHex); err != nil {
		return err
	}
	return nil
}

func (in *AuthorizeInput) Address() common.Address {
	return common.NewAddressFromHex(in.AddressHex)
}

func (in *AuthorizeInput) Signature() common.Signature {
	return common.NewSignatureFromHex(in.SigHex)
}
