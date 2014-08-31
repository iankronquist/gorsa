package main

//import "flag"
import "os"
import "fmt"
import "encoding/json"
import "math/big"


func main() {

}

type keychain_as_json struct {
	priv_exponent string `json:"priv_exponent"`
	priv_modulus string `json:"priv_modulus"`
	pub_exponent string `json:"pub_exponent"`
	pub_modulus string `json:"pub_modulus"`
}

type keychain struct {
	priv_exponent *big.Int
	priv_modulus *big.Int
	pub_exponent *big.Int
	pub_modulus *big.Int
}


func Rsa(base, exponent, modulus *big.Int) *big.Int {
	a := big.NewInt(0)
	return a.Exp(base, exponent, modulus)
}


func Blockify(input []byte, size int, out chan []byte) {
	padding := []byte{}
	for i := 0; i < size - (len(input)  % size); i++ {
		padding = append(padding, byte(0))
	}
	last := 0
	for i := 0; i < (len(input) - 1)/size; i++ {
		begin := i * size;
		end := (i + 1) * size
		last = end
		out <- input[begin:end]
	}
	out <- input[last:]
	out <- padding
	close(out)
}

/*func ReadBytes(file_name string) []byte {

}

func WriteBytes(file_name string, output []byte) {

}*/

func ReadConfig(file_name string) *keychain {
	keys_json := new(keychain_as_json)
	keys := new(keychain)

	file, err := os.Open("keys.json")
	if err != nil {
		fmt.Println("Error opening the file keys.json")
		panic(err)
	}
	file_data := make
	n, err := file.Read(buf)
	if err != nil && err != io.EOF{
		fmt.Println("Error reading the config file keys.json")
		panic(err)
	}
	err = json.Unmarshal(file_data, keys)
	if err != nil {
		fmt.Println("Error reading keys.json")
		panic(err)
	}
	keys.priv_modulus.SetBytes([]byte(keys_json.priv_modulus))
	keys.priv_exponent.SetBytes([]byte(keys_json.priv_exponent))
	keys.pub_modulus.SetBytes([]byte(keys_json.pub_modulus))
	keys.pub_exponent.SetBytes([]byte(keys_json.pub_exponent))
	return keys
}

func Encrypt(message []byte, keys keychain) []byte {
	block_size := 2048
	receive := make(chan []byte)
	output := []byte{}
	go Blockify(message, block_size, receive)
	for block := range receive {
		big_block := big.NewInt(0)
		big_block.SetBytes(block)
		output = append(output, Rsa(big_block, keys.pub_exponent, keys.pub_modulus).Bytes()...)
	}
	return output
}

func Decrypt(message []byte, keys keychain) []byte {
	//block_size := 2048
	//receive := make(chan []byte)
	output := []byte{}
	/*go Blockify(message, block_size, receive)
	for block := range receive {
		output = append(output, Rsa(block, keys.priv_exponent, keys.pub_modulus)...)
	}*/
	return output
}

func Sign(message []byte, keys keychain) []byte {
	return []byte{0}
}

func Verify(message []byte, keys keychain) bool {
	return false
}
