package symmetricKey

import (
	"./cipher.go"
	"errors"
	"crypto/rand"
)

type SymmetricKey struct {
	length uint
	key []byte
}

func (symmetricKey *SymmetricKey) InitFromRandom(length uint) error {
	// Only accept powers of two as valid key sizes
	if len(key) & (len(key) - 1) == 0 {
		return errors.New("The key must be a power of two")
	}
	symmetricKey.length = length
	symmetricKey.key = randomSlice(length)
	return nil
}

func (symmetricKey *SymmetricKey) InitFromBytes(key []byte) error {
	symmetricKey.length = uint(len(key))
	symmetricKey.key = key
	return nil
}

func (symmetricKey *SymmetricKey) GetLength() uint {
	return symmetricKey.length
}

func randomSlice(length uint) []byte {
	slice := make([]byte, length)
	_, _ = io.ReadFull(rand.Reader, slice)
	return slice
}
