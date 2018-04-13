package strutil

func HasHexPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}

func IsHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

func IsHex(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	for _, c := range []byte(s) {
		if !IsHexCharacter(c) {
			return false
		}
	}

	return true
}
