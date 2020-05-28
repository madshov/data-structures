package hashtable

import (
	"fmt"
	"math"
	"math/bits"
	"math/rand"
	"time"
)

// DivHash defines the division hashing method. This is probably the simplest
// hashing method available. A given key is converted into a radix-128 integer,
// and then the remainder is found from dividing the key by the upper bound,
// i.e. number of slots available in the table. A good choice of the upper bound
// will be a prime not too close to power of 2. The division method is heuristic
// in nature, and can succumb to an adversary purposely selecting keys that all
// hash to the same slot, thereby yielding an average retrival time of Θ(n).
type DivHash struct {
	ub int
}

// Hash calculates a hash value using the division method.
func (dh DivHash) Hash(key string) int {
	var (
		k    uint
		hash int
	)

	rs := []rune(key)

	// convert key to a radix-128 integer
	for i := 0; i < len(rs); i++ {
		k += 128*k + uint(rs[i])
	}

	hash = int(k % uint(dh.ub))
	return hash
}

// MulHash defines the multiplication hashing method. This is a simple hashing
// method, where the value of the upper bound, i.e. the number of available
// slots in the table is not critical. A typical choice of the ubber bound is a
// power of 2. The method works by first creating a radix-128 integer of a given
// key, then extracting the fractional part of this multiplied by a constant,
// and finally multiplying the fraction by the upper bound. As the division
// method, this method is also heuristic in nature, so it possesses the same
// issues.
type MulHash struct {
	ub int
}

// Hash calculates a hash value using the multiplication method.
func (mh MulHash) Hash(key string) int {
	var (
		A    float64
		k    uint
		hash int
	)

	rs := []rune(key)

	// convert key to a radix-128 integer
	for i := 0; i < len(rs); i++ {
		k += 128*k + uint(rs[i])
	}

	// This is a good choice for A (see Knuth - Sorting and Searching)
	A = math.Sqrt(5.0) / 2

	// Extract the fractional part of kA, i.e. ⌊kA⌋
	_, f := math.Modf(float64(k) * A)

	// Multiply with upper bound to get hash index
	hash = int(math.Floor(float64(mh.ub) * f))

	return hash
}

// UniHash defines a universal hashing method. This family of hashing methods
// does not suffer the same heuristic behaviour as the division and
// multiplication methods as it includes a randomized bit matrix of log2 to the
// upper bound values.
type UniHash struct {
	m []uint
}

// NewUniHash creates a new instance of UniHash and initializes the randomized
// bit matrix.
func NewUniHash(ub int) UniHash {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	t := int(math.Ceil(math.Log2(float64(ub))))

	m := make([]uint, t)
	for i := 0; i < t; i++ {
		m[i] = uint(r.Int())
	}

	return UniHash{
		m: m,
	}
}

// Hash calculates a hash value using the universal hashing method.
func (uh UniHash) Hash(key string) int {
	var (
		k    uint
		hash int
	)

	rs := []rune(key)

	// convert key to a radix-128 integer
	for i := 0; i < len(rs); i++ {
		k += 128*k + uint(rs[i])
	}

	for i := 0; i < len(uh.m); i++ {
		// find the bit parity
		par := bits.OnesCount(k&uh.m[i]) % 2
		hash <<= 1
		hash |= par
	}
	fmt.Println(hash)
	return hash
}
