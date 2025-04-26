package algebraic

import (
	"errors"
	"math"
)

// Various errors a vector function can return.
var (
	ErrInvalidDim       = errors.New("vector dimension cannot be negative")
	ErrMagZero          = errors.New("vector magnitude cannot be zero")
	ErrIndivisbleByZero = errors.New("vector coordinate is not divisble by zero")
)

// Vector defines a vector structure with a slice of floating point coordinates.
type Vector []float64

// NewVector creates a new instance of a Vector with a given dimension and a
// slice of coordinates, and returns a pointer to it. If the dimension is
// greater than the number of coordinates, the remaining indices of the vector
// will be zero-filled. If the dimension is less, the remaining coordinates will
// be ignored.
func NewVector(dim uint, coords ...float64) Vector {
	cs := make([]float64, dim)
	copy(cs, coords)
	return Vector(cs)
}

// NewUnitVector creates a new instance of a unit vector with a given
// dimension.
func NewUnitVector(dim uint, el uint) Vector {
	cs := make([]float64, dim)
	if dim <= el {
		return Vector(cs)
	}

	cs[el] = 1
	return Vector(cs)
}

// NewZeroVector creates a new instance of a zero-filled vector with a given
// dimension.
func NewZeroVector(dim uint) Vector {
	cs := make([]float64, dim)
	return Vector(cs)
}

// Dimension returns the number of coordinates for the vector.
func (v Vector) Dimension() uint {
	return uint(len(v))
}

// Magnitude returns the distance from the endpoint to the origin for the
// vector.
func (v Vector) Magnitude() float64 {
	var l float64
	for _, c := range v {
		l += c * c
	}

	if l == 0 {
		return 0
	}

	return math.Sqrt(l)
}

// Normalize normalizes, i.e. divides each coodinate with its magnitude for
// the vector.
func (v Vector) Normalize() error {
	mag := v.Magnitude()
	if mag == 0 {
		return ErrMagZero
	}

	for k := range v {
		v[k] /= mag
	}

	return nil
}

// Add adds vector w to vector v.
func (v Vector) Add(w Vector) {
	for k := range v {
		if k < len(w) {
			v[k] += w[k]
		}
	}
}

// Sub subtracts two vectors and returns the resulting vector.
func (v Vector) Sub(w Vector) Vector {
	vec := NewZeroVector(v.Dimension())
	for k, c := range v {
		vec[k] = c - w[k]
	}

	return vec
}

// Mul multiplies two vectors and returns the resulting vector.
func (v Vector) Mul(w Vector) Vector {
	vec := NewZeroVector(v.Dimension())
	for k, c := range v {
		vec[k] = c * w[k]
	}

	return vec
}

// Dic divides two vectors and returns the resulting vector.
func (v Vector) Div(w Vector) (Vector, error) {
	vec := NewZeroVector(v.Dimension())
	for k, c := range v {
		if w[k] == 0 {
			return nil, ErrIndivisbleByZero
		}

		vec[k] = c / w[k]
	}

	return vec, nil
}

// Dot returns the dot product (scalar product) of two vectors.
func (v Vector) Dot(w Vector) float64 {
	var dot float64
	for k, c := range v {
		dot += c * w[k]
	}

	return dot
}

// Scale scales each coordinate in the vector with a given scalar value.
func (v Vector) Scale(scalar float64) Vector {
	vec := NewZeroVector(v.Dimension())
	for k, c := range v {
		vec[k] = c * scalar
	}

	return vec
}

// GetCoord returns a given coordinate for the vector.
func (v Vector) GetCoord(coord uint) (float64, error) {
	if v.Dimension() < coord+1 {
		return 0, ErrInvalidDim
	}

	return v[coord], nil
}

// X returns the first coordinate of the vector. Shorthand for v.coords[0].
func (v Vector) X() float64 {
	c, _ := v.GetCoord(0)
	return c
}

// Y returns the second coordinate of the vector. Shorthand for v.coords[1].
func (v Vector) Y() float64 {
	c, _ := v.GetCoord(1)
	return c
}

// Z returns the third coordinate of the vector. Shorthand for v.coords[2].
func (v Vector) Z() float64 {
	c, _ := v.GetCoord(2)
	return c
}
