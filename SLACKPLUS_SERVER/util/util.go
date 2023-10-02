package util

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
)

var alphabet string = "abcdefghijklmnopqrstuvwxyz"

func Encrypt(str string) string {
	encrypter := sha256.New()
	encrypter.Write([]byte(str))
	encryptedBS := encrypter.Sum(nil)

	return fmt.Sprintf("%x", encryptedBS)
}

func RandomString() string {
	var builder strings.Builder

	for i := 0; i < 26; i++ {
		builder.WriteString(string(alphabet[rand.Intn(26)]))
	}

	return builder.String()
}
