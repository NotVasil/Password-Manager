package password

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Encrypt(value string, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}

	plaintext := []byte(value)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	return Encode(ciphertext)
}
