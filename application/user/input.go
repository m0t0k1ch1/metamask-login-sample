package user

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/user"
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
	if err := common.ValidateAddressHex(in.AddressHex); err != nil {
		return err
	}
	return nil
}

func (in *GetUserInput) Address() common.Address {
	return common.NewAddressFromHex(in.AddressHex)
}

type UpdateUserInput struct {
	AddressHex string
	Name       string
}

func NewUpdateUserInput(addressHex, name string) *UpdateUserInput {
	return &UpdateUserInput{
		AddressHex: addressHex,
		Name:       name,
	}
}

func (in *UpdateUserInput) Validate() error {
	if err := common.ValidateAddressHex(in.AddressHex); err != nil {
		return nil
	}
	if err := user.ValidateUserName(in.Name); err != nil {
		return err
	}
	return nil
}

func (in *UpdateUserInput) Address() common.Address {
	return common.NewAddressFromHex(in.AddressHex)
}
