package monokage

type DB struct {
	bf *bloomFilter
}

func (db *DB) Insert(key []byte) {
	db.bf.Add(key)
}

func (db *DB) Check(key []byte) bool {
	return db.bf.Test(key)
}

func NewDB(m uint, k uint) *DB {
	return &DB{
		bf: newBloomFilter(m, k),
	}
}
