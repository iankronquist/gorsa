package rsa

import (
	"math/big"
	"../key/rsakey"
)

func RSA (message, priv_key, modulus *big.Int) *big.Int {
	return message.Exp(message, priv_key, modulus)
}

func RSA_Bytes (message [] byte, key, modulus *big.Int) [] byte {
	message_num := big.Int{}
	message_num.SetBytes(message)
	cipher_text_num := RSA(&message_num, key, modulus)
	return cipher_text_num.Bytes()
}

func Encrypt
