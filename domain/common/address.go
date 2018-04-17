package common

import "github.com/ethereum/go-ethereum/common"

type Address common.Address

func (address Address) Hex() string {
	return common.Address(address).Hex()
}

func NewAddressFromHex(addressHex string) Address {
	return Address(common.HexToAddress(addressHex))
}

func ValidateAddressHex(addressHex string) error {
	if !common.IsHexAddress(addressHex) {
		return ErrInvalidAddressHex
	}
	return nil
}
