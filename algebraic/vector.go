package algebraic

import (
	"errors"
	"math"
)

// Various errors a vector function can return.
var (
	ErrInvalidDim       = errors.New("vector dimension cannot be negative")
	ErrCoordOutOfBounds = errors.New("vector coordinate is greater than dimension")
	ErrMagZero          = errors.New("vector magnitude cannot be zero")
	ErrDimsNotEqual     = errors.New("vector dimensions do not match")
	ErrIndivisbleByZero = errors.New("vector coordinate is not divisble by zero")
)

// Vector defines a vector structure with a slice of floating point coordinates.
type Vector []float64

// NewVector creates a new instance of a Vector with a given dimension and a
// slice of coordinates, and returns a pointer to it. If the dimension is
// greater than the number of coordinates, the remaining indices of the vector
// will be zero-filled. If the dimension is less, the remaining coordinates will
// be ignored.
func NewVector(dim int, coords ...float64) (Vector, error) {
	if dim < 0 {
		return nil, ErrInvalidDim
	}

	cs := make([]float64, dim)
	copy(cs, coords)
	return Vector(cs), nil
}

// NewZeroVector creates a new instance of a zero-filled vector with a given
// dimension.
func NewZeroVector(dim int) (Vector, error) {
	return NewVector(dim)
}

// NewUnitVector creates a new instance of a unit vector with a given
// dimension.
func NewUnitVector(dim, coord int) (Vector, error) {
	v, err := NewVector(dim)
	if err != nil {
		return nil, err
	}

	if len(v) <= coord {
		return nil, ErrCoordOutOfBounds
	}

	v[coord] = 1
	return v, nil
}

// Dimension returns the number of coordinates for the vector.
func (v Vector) Dimension() int {
	return len(v)
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

// Add adds two vectors and returns the resulting vector.
func (v Vector) Add(w Vector) (Vector, error) {
	vec, err := NewZeroVector(v.Dimension())
	if err != nil {
		return nil, err
	}

	for k, c := range v {
		vec[k] = c + w[k]
	}

	return vec, nil
}

// Sub subtracts two vectors and returns the resulting vector.
func (v Vector) Sub(w Vector) (Vector, error) {
	vec, err := NewZeroVector(v.Dimension())
	if err != nil {
		return nil, err
	}

	for k, c := range v {
		vec[k] = c - w[k]
	}

	return vec, nil
}

// Mul multiplies two vectors and returns the resulting vector.
func (v Vector) Mul(w Vector) (Vector, error) {
	vec, err := NewZeroVector(v.Dimension())
	if err != nil {
		return nil, err
	}

	for k, c := range v {
		vec[k] = c * w[k]
	}

	return vec, nil
}

// Dic divides two vectors and returns the resulting vector.
func (v Vector) Div(w Vector) (Vector, error) {
	vec, err := NewZeroVector(v.Dimension())
	if err != nil {
		return nil, err
	}

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
func (v Vector) Scale(scalar float64) (Vector, error) {
	vec, err := NewZeroVector(v.Dimension())
	if err != nil {
		return nil, err
	}

	for k, c := range v {
		vec[k] = c * scalar
	}

	return vec, nil
}

// GetCoord returns a given coordinate for the vector.
func (v Vector) GetCoord(coord int) float64 {
	if v.Dimension() < coord+1 {
		return 0
	}

	return v[coord]
}

// X returns the first coordinate of the vector. Shorthand for v.coords[0].
func (v Vector) X() float64 {
	return v.GetCoord(0)
}

// Y returns the second coordinate of the vector. Shorthand for v.coords[1].
func (v Vector) Y() float64 {
	return v.GetCoord(1)
}

// Z returns the third coordinate of the vector. Shorthand for v.coords[2].
func (v Vector) Z() float64 {
	return v.GetCoord(2)
}
