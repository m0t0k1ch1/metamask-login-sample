package user

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/spec"
)

type GetUserInput struct {
	AddressHex string
}

func NewGetUserInput(addressHex string) *GetUserInput {
	return &GetUserInput{
		AddressHex: addressHex,
	}
}

func (in *GetUserInput) Validate() error {
	if ok := spec.IsValidAddressHex(in.AddressHex); !ok {
		return domain.ErrInvalidAddressHex
	}
	return nil
}

func (in *GetUserInput) Address() model.Address {
	return model.NewAddressFromHex(in.AddressHex)
}
