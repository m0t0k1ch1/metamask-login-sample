package application

import "github.com/ethereum/go-ethereum/common"

func hexToAddress(addressHex string) (common.Address, error) {
	if !common.IsHexAddress(addressHex) {
		// TODO: return error
		return common.Address{}, nil
	}

	return common.HexToAddress(addressHex), nil
}
