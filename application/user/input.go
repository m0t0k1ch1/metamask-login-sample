package user

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
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
	if ok := common.IsValidAddressHex(in.AddressHex); !ok {
		return domain.ErrInvalidAddressHex
	}
	return nil
}

func (in *GetUserInput) Address() common.Address {
	return common.NewAddressFromHex(in.AddressHex)
}
