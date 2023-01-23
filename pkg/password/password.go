package password

import (
	"crypto/rand"
	"unsafe"
)

var characters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#")
var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func RandomPassword(size int) string {
	bytes := make([]byte, size)
	rand.Read(bytes)

	for i := 0; i < size; i++ {
		bytes[i] = characters[bytes[i]%byte(len(characters))]
	}

	return *(*string)(unsafe.Pointer(&bytes))
}
