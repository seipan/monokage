package monokage

import "github.com/bits-and-blooms/bitset"

type BloomFilter struct {
	m uint
	k uint
	b *bitset.BitSet
}
