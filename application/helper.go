package application

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
)

func hexToAddress(addressHex string) (common.Address, error) {
	if !common.IsHexAddress(addressHex) {
		return common.Address{}, model.ErrInvalidAddress
	}

	return common.HexToAddress(addressHex), nil
}
