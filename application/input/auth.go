package input

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/spec"
)

type AuthChallengeInput struct {
	AddressHex string
}

func NewAuthChallengeInput(addressHex string) *AuthChallengeInput {
	return &AuthChallengeInput{
		AddressHex: addressHex,
	}
}

func (input *AuthChallengeInput) Validate() error {
	if ok := spec.IsValidHexAddress(input.AddressHex); !ok {
		return domain.ErrInvalidAddressFormat
	}
	return nil
}

func (input *AuthChallengeInput) Address() common.Address {
	return common.HexToAddress(input.AddressHex)
}

type AuthAuthorizeInput struct {
	AddressHex string
	SigHex     string
}

func NewAuthAuthorizeInput(addressHex, sigHex string) *AuthAuthorizeInput {
	return &AuthAuthorizeInput{
		AddressHex: addressHex,
		SigHex:     sigHex,
	}
}

func (input *AuthAuthorizeInput) Validate() error {
	if ok := spec.IsValidHexAddress(input.AddressHex); !ok {
		return domain.ErrInvalidAddressFormat
	}
	if ok := spec.IsValidHexSignature(input.SigHex); !ok {
		return domain.ErrInvalidSignatureFormat
	}
	return nil
}

func (input *AuthAuthorizeInput) Address() common.Address {
	return common.HexToAddress(input.AddressHex)
}

func (input *AuthAuthorizeInput) SigBytes() []byte {
	return common.FromHex(input.SigHex)
}
