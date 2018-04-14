package spec

import "github.com/ethereum/go-ethereum/common"

func IsValidAddressHex(addressHex string) bool {
	return common.IsHexAddress(addressHex)
}
