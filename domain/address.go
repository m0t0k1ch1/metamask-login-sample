package domain

import "github.com/ethereum/go-ethereum/common"

type Address common.Address

func NewAddressFromHex(addressHex string) Address {
	return Address(common.HexToAddress(addressHex))
}

func (address Address) Hex() string {
	return common.Address(address).Hex()
}

func ValidateAddressHex(addressHex string) error {
	if !common.IsHexAddress(addressHex) {
		return ErrInvalidAddressHex
	}
	return nil
}
