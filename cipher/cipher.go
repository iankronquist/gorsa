package cipher

import (
	"math/big"
	"../key"
)

type Cipher interface {
	Encrypt(message *big.Int, cipherKey *key.Key) *big.Int
	Decrypt(message *big.Int, cipherKey *key.Key) *big.Int
	Encrypt_Bytes(message *[]byte, cipherKey *key.Key) *[]byte
	Decrypt_Bytes(cipherText *[]byte, cipherKey *key.Key) *[]byte
}

