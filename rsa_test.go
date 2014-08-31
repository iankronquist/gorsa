package main

import "testing"
import "math/big"
import "bytes"

func Test_Rsa(t *testing.T) {
	e := big.NewInt(17)
	d := big.NewInt(2753)
	n := big.NewInt(3233)
	m := big.NewInt(65)
	c := big.NewInt(2790)
	cipher_output := Rsa(m, e, n)
	if (cipher_output.Cmp(c) != 0) {
		t.Error("Encryption failed. Got ", cipher_output, " expected ", c, )
	}
	plain_output := Rsa(cipher_output, d, n)
	if (plain_output.Cmp(m) != 0) {
		t.Error("Decryption failed. Got ", plain_output,
			" expected ", m)
	}
}

func Test_To_Block(t *testing.T) {
	s := []byte("In fact, the chain doesn’t fall off unless the bent spoke and the weak link happen to coincide.")
	s_received := []byte("")
	s_expected := []byte("In fact, the chain doesn’t fall off unless the bent spoke and the weak link happen to coincide.xxx")
	block_size := 10
	times_iter := 0
	times_iter_expected := (len(s_expected))/block_size + 1
	receive := make(chan []byte)
	go Blockify(s, block_size, receive)
	for substr := range receive {
		//s_received += substr
		s_received = append(s_received, substr...)
		times_iter += 1
	}
	if (times_iter != times_iter_expected) {
		t.Error("Didn't make the right number of blocks. Made", times_iter,
			" expected ", times_iter_expected)
	}
	if bytes.Equal(s_received, s_expected) {
		t.Error("The received string was no what was expected. Expected\n",
			s_expected, "\ngot\n", s_received)
	}
}

func Test_Encrypt_and_Decrypt(t *testing.T) {
	s := []byte("\"Talk, talk, talk!\" says Alan Turing, imitating the squawk of furious hens. The strange noise is made stranger by the fact that he is wearing a gas mask, until he becomes impatient and pulls it up onto his forehead. They love to hear themselves talk.")
	k := new(keychain)
	k.pub_exponent = big.NewInt(17)
	k.priv_exponent = big.NewInt(2753)
	k.pub_modulus = big.NewInt(3233)
	k.priv_modulus = big.NewInt(3120)

	cipher_text := Encrypt(s, *k)
	if bytes.Equal(cipher_text, s) {
		t.Error("Cipher text should not equal message.")
	}
	plain_text := Decrypt(cipher_text, *k)
	if !bytes.Equal(plain_text, s) {
		t.Error("Plain text must equal original message")
	}
}

func Test_Sign_and_Verify(t *testing.T) {
	s := []byte("\"Talk, talk, talk!\" says Alan Turing, imitating the squawk of furious hens. The strange noise is made stranger by the fact that he is wearing a gas mask, until he becomes impatient and pulls it up onto his forehead. They love to hear themselves talk.")
	k := new(keychain)
	k.pub_exponent = big.NewInt(17)
	k.priv_exponent = big.NewInt(2753)
	k.pub_modulus = big.NewInt(3233)
	k.priv_modulus = big.NewInt(3120)

	signature := Sign(s, *k)
	if !Verify(signature, *k){
		t.Error("Plain text must equal original message")
	}
}
