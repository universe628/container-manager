package utils

import "encoding/hex"

func ToHex(str []byte) string {
	return hex.EncodeToString(str)
}
