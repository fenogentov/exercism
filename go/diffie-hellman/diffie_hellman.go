package diffiehellman

import (
	rand "crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) (a *big.Int) {
	a, err := rand.Int(rand.Reader, p)
	if err != nil {
		panic(err)
	}
	for a.Cmp(big.NewInt(1)) < 1 {
		a.Add(a, big.NewInt(1))
	}
	return a
}

func PublicKey(a, p *big.Int, g int64) (A *big.Int) {
	A = new(big.Int).Exp(big.NewInt(g), a, p)
	return A
}

func SecretKey(a, B, p *big.Int) (s *big.Int) {
	s = new(big.Int).Exp(B, a, p)
	return s
}

func NewPair(p *big.Int, g int64) (a, A *big.Int) {
	a = PrivateKey(p)
	A = PublicKey(a, p, g)
	return a, A
}
