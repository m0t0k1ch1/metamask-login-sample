package strutil

import (
	"math/rand"
	"time"
)

const (
	letters        = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lettersIdxMask = 1<<6 - 1
)

func Rand(n int) string {
	randSrc := rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	for i := 0; i < n; {
		idx := int(randSrc.Int63() & lettersIdxMask)
		if idx < len(letters) {
			b[i] = letters[idx]
			i++
		}
	}

	return string(b)
}
