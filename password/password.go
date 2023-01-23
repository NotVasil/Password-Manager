package password

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"log"
	"unsafe"
)

var characters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#")
var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	return data
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

func RandomPassword(size int) string {
	bytes := make([]byte, size)
	rand.Read(bytes)

	for i := 0; i < size; i++ {
		bytes[i] = characters[bytes[i]%byte(len(characters))]
	}

	return *(*string)(unsafe.Pointer(&bytes))
}
