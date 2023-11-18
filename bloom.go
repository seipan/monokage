package monokage

import (
	"github.com/bits-and-blooms/bitset"
)

type bloomFilter struct {
	m uint
	k uint
	b *bitset.BitSet
}
