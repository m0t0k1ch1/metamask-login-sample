package application

import "github.com/m0t0k1ch1/metamask-login-sample/domain/common"

type AddressHexInput struct {
	AddressHex string
}

func NewAddressHexInput(addressHex string) *AddressHexInput {
	return &AddressHexInput{
		AddressHex: addressHex,
	}
}

func (in *AddressHexInput) Validate() error {
	if err := common.ValidateAddressHex(in.AddressHex); err != nil {
		return err
	}
	return nil
}

func (in *AddressHexInput) Address() common.Address {
	return common.NewAddressFromHex(in.AddressHex)
}
