package monokage

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)

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

func getMD5Hash(str []byte) string {
	hasher := md5.New()
	hasher.Write(str)

	return hex.EncodeToString(hasher.Sum(nil))
}
