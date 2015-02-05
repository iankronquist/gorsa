package key

// Define the key interface
type Key interface {
	GetLength() uint
}

// Define interfaces for symmetric and asymmetric keys
type SymmetricKey struct {
	length uint
	key []byte
}

type AsymmetricKey struct {
	length uint
	pub_key, priv_key, modulus []byte
}
