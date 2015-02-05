package rsa

import (
	"math/big"
	"../key/rsakey"
)

type RSACipher struct {}

func (cipher *RSACipher) RSA (message, priv_key, modulus *big.Int) *big.Int {
	return message.Exp(message, priv_key, modulus)
}

func (cipher *RSACipher) RSA_Bytes (message [] byte, key, modulus *big.Int) [] byte {
	message_num := big.Int{}
	message_num.SetBytes(message)
	cipher_text_num := RSA(&message_num, key, modulus)
	return cipher_text_num.Bytes()
}

func (cipher *RSACipher) Encrypt_Bytes (message [] byte, key,
	modulus *big.Int) [] byte {
	return RSA_Bytes(message, key, modulus)
}

func (cipher *RSACipher) Decrypt_Bytes (message [] byte, key,
	modulus *big.Int) [] byte {
	return RSA_Bytes(message, key, modulus)
}

func (cipher *RSACipher) Encrypt (message, pub_key, modulus *big.Int) *big.Int {
	RSA (message, pub_key, modulus)
}

func (cipher *RSACipher) Decrypt (message, priv_key, modulus *big.Int) *big.Int {
	RSA (message, priv_key, modulus)
}
