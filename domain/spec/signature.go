package spec

import (
	"github.com/m0t0k1ch1/metamask-login-sample/domain/model"
	"github.com/m0t0k1ch1/metamask-login-sample/library/strutil"
)

func IsValidSignatureHex(sigHex string) bool {
	if strutil.HasHexPrefix(sigHex) {
		sigHex = sigHex[2:]
	}

	return len(sigHex) == 2*model.SignatureLength && strutil.IsHex(sigHex)
}
