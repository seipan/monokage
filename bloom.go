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

func (bf *bloomFilter) Test(key []byte) bool {
	hash := getMD5Hash(key)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64_hashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64_hashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := uint(0); i < bf.k; i++ {
		if !bf.b.Test(uint(doubleHashing(int64(i64_hashA), int64(i64_hashB), int(i), int(bf.m)))) {
			return false
		}
	}
	return true
}

func (bf *bloomFilter) TestAndAdd(key []byte) bool {
	var result bool
	result = true
	hash := getMD5Hash(key)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64_hashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64_hashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := uint(0); i < bf.k; i++ {
		if bf.b.Test(uint(doubleHashing(int64(i64_hashA), int64(i64_hashB), int(i), int(bf.m)))) {
			result = false
		}
		bf.b.Set(uint(doubleHashing(int64(i64_hashA), int64(i64_hashB), int(i), int(bf.m))))
	}
	return result
}

func newBloomFilter(m uint, k uint) *bloomFilter {
	return &bloomFilter{
		m: m,
		k: k,
		b: bitset.New(uint(m)),
	}
}
