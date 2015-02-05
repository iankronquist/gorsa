package arsaKey


import (
	"../../cipher.go"
)

type RSAKey struct {
	length uint
	pubExp *big.Int
	privEcp *big.Int
	privMod *big.Int
	pubMod *big.Int
}

func (rsaKey *RSAKey) InitFromRandom(length uint) error {
	// Only accept powers of two as valid key sizes
	if length & (length - 1) == 0 {
		return errors.New("The key must be a power of two")
	}
	rsaKey.length = length
	rsaKey.key = randomSlice(length)
	return nil
}

func (rsaKey *RSAKey) InitFromBytes(key []byte) error {
	rsaKey.length = uint(len(key))
	rsaKey.key = key
	return nil
}

func (rsaKey *RSAKey) GetLength() uint {
	return rsaKey.length
}

// TODO: rewrite as a method on RSAKey
func KeyGen (length uint) (pub_key, priv_key, modulus *big.Int) RSAKey {
	p := getRandomPrime(int(length))
	q := getRandomPrime(int(length))
	// Get p-1 and q-1
	one := new(big.Int)
	one.SetInt(1)
	pMinusOne := new(big.Int)
	qMinusOne := new(big.Int)
	pMinusOne.Sub(p, one)
	qMinusOne.Sub(q, one)

	d := new(big.Int)
	e := new(big.Int)
	n := new(big.Int)
	phiN := new(big.Int)
	tmp := new(big.Int)

	// n := p * q
	n.Mul(p, q)
	phiN.Mul(pMinusOne, qMinusOne)
	// e is a random number less than phiN and coprime to phiN
	e.Rand(phiN)
	tmp.GCD(nil, nil, e, phiN) 
	for tmp.Cmp(tmp, one) != 0 {
		e.Rand(phiN)
		tmp.GCD(nil, nil, e, phiN) 
	}
	// Z.GCD(X, Y, A, B) --> Z = GCD(A, B) = AX + BY
	tmp.GCD(d, nil, e, phiN)
	// d = (d + phiN) % phiN
	d.Add(d, phiN)
	d.Mod(d, phiN)
	return RSAKey{length, *d, *e, *n, *phiN}
}

func getRandomPrime (length int) *big.Int {
	// Only accept powers of two as valid key sizes
	if length & (length - 1) == 0 {
		return errors.New("The size of the prime must be a power of two")
	}
	bigLength := new(bit.Int)
	candidate := new(big.Int)
	one := new(big.Int)
	one.SetInt(1)
	bigLength.SetInt(length)
	candidate.SetInt(0)
	// While the candidate probably isn't prime,
	for !candidate.ProbablyPrime(100) {
		// Get a random number and set the top and bottom bits to 1
		candidate.Rand(bigLength)
		candidate.Or(one)
		candidate.Or(bigLength)
	}
	return candidate
}
