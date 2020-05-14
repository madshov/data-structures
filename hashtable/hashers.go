package hashtable

import (
	"math"
)

// DivHash defines the division hashing method.
type DivHash struct{}

// Hash calculates a hash value using the division method. This is probably
// the simplest hashing method available. The key is converted into a radix-128
// integer, and then the remainder is found from dividing the key by the upper
// bound, i.e. number of slots available. A good choice of the upper bound will
// be a prime not too close to power of 2. The division method is heuristic in
// nature, and can succumb to an adversary purposely selecting keys that all
// hash to the same slot, thereby yielding an average retrival time of Θ(n).
func (dh DivHash) Hash(lb, ub int, key string) int {
	var k int

	rs := []rune(key)

	// convert key to a radix-128 integer
	for i := 0; i < len(rs); i++ {
		k += 128*k + int(rs[i])
	}

	return k % ub
}

// MulHash defines the multiplication hashing method.
type MulHash struct{}

// Hash calculates a hash value using the multiplication method. This is a
// simple hashing method, where the value of the upper bound, i.e. the number of
// available slots is not critical. A typical choice of the ubber bound is a
// power of 2. The method works by first creating a radix-128 integer of the
// key, then extracting the fractional part of this multiplied by a constant,
// and finally multiplying the fraction by the upper bound. As the division
// method, this method is also heuristic in nature, so it possesses the same
// issues.
func (mh MulHash) Hash(lb, ub int, key string) int {
	var (
		A float64
		k int
	)

	rs := []rune(key)

	// convert key to a radix-128 integer
	for i := 0; i < len(rs); i++ {
		k += 128*k + int(rs[i])
	}

	// This is a good choice for A (see Knuth - Sorting and Searching)
	A = math.Sqrt(5.0) / 2

	// Extract the fractional part of kA, i.e. ⌊kA⌋
	_, f := math.Modf(float64(k) * A)

	// Multiply with upper bound to get hash index
	return int(math.Floor(float64(ub) * f))
}
