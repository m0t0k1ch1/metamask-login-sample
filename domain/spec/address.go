package spec

import "github.com/ethereum/go-ethereum/common"

func IsValidHexAddress(addressHex string) bool {
	return common.IsHexAddress(addressHex)
}
