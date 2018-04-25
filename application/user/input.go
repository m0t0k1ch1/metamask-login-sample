package user

import (
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/domain"
)

type GetUserInput struct {
	*application.AddressHexInput
}

func NewGetUserInput(addressHex string) *GetUserInput {
	return &GetUserInput{
		AddressHexInput: application.NewAddressHexInput(addressHex),
	}
}

type UpdateUserInput struct {
	*application.AddressHexInput
	Name string
}

func NewUpdateUserInput(addressHex, name string) *UpdateUserInput {
	return &UpdateUserInput{
		AddressHexInput: application.NewAddressHexInput(addressHex),
		Name:            name,
	}
}

func (in *UpdateUserInput) Validate() error {
	if err := in.AddressHexInput.Validate(); err != nil {
		return err
	}
	if err := domain.ValidateUserName(in.Name); err != nil {
		return err
	}
	return nil
}

type DeleteUserInput struct {
	*application.AddressHexInput
}

func NewDeleteUserInput(addressHex string) *DeleteUserInput {
	return &DeleteUserInput{
		AddressHexInput: application.NewAddressHexInput(addressHex),
	}
}
