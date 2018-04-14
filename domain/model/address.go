package model

import "github.com/ethereum/go-ethereum/common"

type Address common.Address

func (address Address) Hex() string {
	return common.Address(address).Hex()
}

func NewAddressFromHex(addressHex string) Address {
	return Address(common.HexToAddress(addressHex))
}
