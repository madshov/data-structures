package hashtable

import (
	"errors"

	"github.com/madshov/data-structures/elementary"
)

// Various errors a list function can return.
var (
	ErrNonExistingKey = errors.New("object does not exist for key")
	ErrInternal       = errors.New("internal error")
)

// New creates a new hash table with a given size and hashing mechanism.
func New(s int, h Hasher) *HashTable {
	ls := make([]*elementary.List, s)

	for i := 0; i < s; i++ {
		l := elementary.NewLinkedList()
		ls[i] = l
	}

	// set default hashing method
	if h == nil {
		h = MulHash{}
	}

	return &HashTable{
		list:   ls,
		hasher: h,
		size:   s,
	}
}

// HashTable defines a hash table structure with a slice of linked lists, a
// size, and a hashing mechanism,
type HashTable struct {
	list   []*elementary.List
	hasher Hasher
	size   int
}

// Insert adds a given key and value pair to the hash table.
func (ht *HashTable) Insert(key string, val int) {
	idx := ht.hasher.Hash(0, ht.size, key)
	ht.list[idx].Insert(key, val)
}

// Search looks for an object with a given key and, if it exists, returns its
// value in the hash table.
func (ht *HashTable) Search(key string) (int, bool) {
	idx := ht.hasher.Hash(0, ht.size, key)
	obj, exists := ht.list[idx].Search(key)
	if !exists {
		return 0, false
	}

	return obj.Value, true
}

// Delete looks for an object with a given key and, if it exists, deletes it
// from the hash table.
func (ht *HashTable) Delete(key string) error {
	idx := ht.hasher.Hash(0, ht.size, key)
	obj, exists := ht.list[idx].Search(key)
	if !exists {
		return ErrNonExistingKey
	}

	if err := ht.list[idx].Delete(obj); err != nil {
		return ErrInternal
	}

	return nil
}

// Hasher is the hasher used for the hash table operations.
type Hasher interface {
	Hash(int, int, string) int
}
