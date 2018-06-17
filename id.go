package id

// Id package generates true random byte sequence,
// using "crypto/rand" package, and it also can encode it to hex.
// You have to set a number of bytes, which should be generated.
// After hex encoding, the number of the symbols is doubled.


import (
	"log"
	"errors"
	"crypto/rand"
	"encoding/hex"
)

func GetRandomBytes(nBytes int) ([]byte, error) {
	randoms := make([]byte, nBytes)
	n, err := rand.Read(randoms)
	if err != nil {
		return nil, err
	}
	if n != nBytes {
		return nil, errors.New("Can't generate required number of bytes")
	}
	return randoms, nil
}

func toHex(b []byte) []byte {
	hexed := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(hexed, b)
	return hexed
}

func GetRandomHex(n int) ([]byte, error) {
	randoms, err := GetRandomBytes(n)
	if err != nil {
		return nil, err
	}
	hexed := toHex(randoms)
	return hexed, nil
}

func GetRandomHexString(n int) (string, error) {
	randoms, err := GetRandomHex(n)
	if err != nil {
		return "", err
	}
	return string(randoms), nil
}
