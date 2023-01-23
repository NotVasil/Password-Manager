package password

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

func Decrypt(value string, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}

	ciphertext := Decode(value)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plainText, ciphertext)

	return string(plainText)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	return data
}
