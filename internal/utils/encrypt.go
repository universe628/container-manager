package utils

import (
	"crypto/sha512"
)

func Sha512Hash(str string) []byte {
	hash := sha512.New()
	hash.Write([]byte(str))
	return hash.Sum(nil)
}
