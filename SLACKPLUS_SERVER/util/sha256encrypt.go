package util

import (
	"crypto/sha256"
	"fmt"
)

func Encrypt(str string) string {
	encrypter := sha256.New()
	encrypter.Write([]byte(str))
	encryptedBS := encrypter.Sum(nil)

	return fmt.Sprintf("%x", encryptedBS)
}
