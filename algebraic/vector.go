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
type vector struct {
	coords []float64
}

// NewVector creates a new instance of a Vector with a given dimension and a
// slice of coordinates, and returns a pointer to it. If the dimension is
// greater than the number of coordinates, the remaining indexes of the vector
// will be zero-filled. If the dimension is less, the remaining coordinates will
// be ignored.
func NewVector(dim int, coords ...float64) (*vector, error) {
	if dim < 0 {
		return nil, ErrInvalidDim
	}

	cs := make([]float64, dim)
	for i := 0; i < dim; i++ {
		cs[i] = coords[i]
	}

	return &vector{
		coords: cs,
	}, nil
}

// NewZeroVector creates a new instance of a zero-filled vector with a given
// dimension.
func NewZeroVector(dim int) (*vector, error) {
	return NewVector(dim)
}

// NewUnitVector creates a new instance of a unit vector with a given
// dimension.
func NewUnitVector(dim, coord int) (*vector, error) {
	vec, err := NewVector(dim)
	if err != nil {
		return nil, err
	}

	if len(vec.coords) < coord {
		return nil, ErrCoordOutOfBounds
	}

	vec.coords[coord] = 1
	return vec, nil
}

// Dimension returns the number of coordinates for the vector.
func (v *vector) Dimension() int {
	return len(v.coords)
}

// Magnitude returns the distance from the endpoint to the origin for the
// vector.
func (v *vector) Magnitude() float64 {
	var l float64
	for _, c := range v.coords {
		l += c * c
	}

	if l == 0 {
		return 0
	}

	return math.Sqrt(l)
}

// Normalize normalizes, i.e. divides each coodinate with its magnitude, a for
// the vector.
func (v *vector) Normalize() error {
	mag := v.Magnitude()
	if mag == 0 {
		return ErrMagZero
	}

	for k := range v.coords {
		v.coords[k] /= mag
	}

	return nil
}

// Add adds each coordinate of two vectors.
func (v *vector) Add(w *vector) *vector {
	rv := &vector{
		coords: make([]float64, v.Dimension()),
	}

	for k, c := range v.coords {
		rv.coords[k] = c + w.coords[k]
	}

	return rv
}

// Sub substracts each coordinate of two vectors.
func (v *vector) Sub(w *vector) *vector {
	rv := &vector{
		coords: make([]float64, v.Dimension()),
	}

	for k, c := range v.coords {
		rv.coords[k] = c - w.coords[k]
	}

	return rv
}

// Mul multiplies each coordinate of two vectors.
func (v *vector) Mul(w *vector) *vector {
	rv := &vector{
		coords: make([]float64, v.Dimension()),
	}

	for k, c := range v.coords {
		rv.coords[k] = c * w.coords[k]
	}

	return rv
}

// Div divides each coordinate of two vectors.
func (v *vector) Div(w *vector) (*vector, error) {
	rv := &vector{
		coords: make([]float64, v.Dimension()),
	}

	for k, c := range v.coords {
		if w.coords[k] == 0 {
			return nil, ErrIndivisbleByZero
		}
		rv.coords[k] = c / w.coords[k]
	}

	return rv, nil
}

// Dot returns the dot product (scalar product) of two vectors.
func (v *vector) Dot(w *vector) float64 {
	var d float64
	for k, c := range v.coords {
		d += c * w.coords[k]
	}

	return d
}

// Scale scales each coordinate in the vector with a given scalar value.
func (v *vector) Scale(scalar float64) *vector {
	rv := &vector{
		coords: make([]float64, v.Dimension()),
	}

	for k, c := range v.coords {
		rv.coords[k] = c * scalar
	}

	return rv
}

// GetCoord returns a given coordinate for the vector.
func (v *vector) GetCoord(coord int) float64 {
	return v.coords[coord]
}
