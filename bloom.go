package monokage

import (
	"strconv"

	"github.com/bits-and-blooms/bitset"
)

type bloomFilter struct {
	m uint
	k uint
	b *bitset.BitSet
}

func (bf *bloomFilter) Add(key []byte) {
	hash := getMD5Hash(key)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64_hashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64_hashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := uint(0); i < bf.k; i++ {
		bf.b.Set(uint(doubleHashing(i64_hashA, i64_hashB, int(i), int(bf.m))))
	}
}

func newBloomFilter(m uint, k uint) *bloomFilter {
	return &bloomFilter{
		m: m,
		k: k,
		b: bitset.New(uint(m)),
	}
}
