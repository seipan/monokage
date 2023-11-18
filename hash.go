package monokage

import "math/big"

func doubleHashing(hashA, hashB int64, n int, size int) (hash int64) {
	h := new(big.Int).Mul(big.NewInt(int64(n)), big.NewInt(hashB))
	h = new(big.Int).Add(big.NewInt(hashA), h)
	h = new(big.Int).Rem(h, big.NewInt(int64(size)))

	hash = h.Int64()
	if hash < 0 {
		hash += int64(size)
	}
	return
}
