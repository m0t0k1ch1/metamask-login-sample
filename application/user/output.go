package user

import "github.com/m0t0k1ch1/metamask-login-sample/domain"

type GetUserOutput struct {
	Name       string `json:"name"`
	AddressHex string `json:"address"`
}

func NewGetUserOutput(u *domain.User) *GetUserOutput {
	return &GetUserOutput{
		Name:       u.Name,
		AddressHex: u.Address.Hex(),
	}
}

type UpdateUserOutput struct{}

func NewUpdateUserOutput() *UpdateUserOutput {
	return &UpdateUserOutput{}
}

type DeleteUserOutput struct{}

func NewDeleteUserOutput() *DeleteUserOutput {
	return &DeleteUserOutput{}
}
