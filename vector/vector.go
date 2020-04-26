package vector

import (
	"errors"
	"math"
)

// NewVector creates a new instance of a Vector with a given slice of
// coordinates, and returns a pointer to it. The dimension is determined by the
// number of coordinates.
func NewVector(cs ...float64) *Vector {
	return &Vector{
		coords: cs,
	}
}

// NewZeroVector creates a new instance of a zero-filled Vector with a given
// dimension .
func NewZeroVector(dim int) *Vector {
	cs := make([]float64, dim)
	return &Vector{
		coords: cs,
	}
}

// Vector defines a vector structure with a slice of floating point coordinates.
type Vector struct {
	coords []float64
}

// Various errors a vector function can return.
var (
	ErrMagZero      = errors.New("vector magnitude cannot be zero")
	ErrDimsNotEqual = errors.New("vector dimensions do not match")
)

// Dimension returns the number of coordinates for a given vector.
func (v *Vector) Dimension() int {
	return len(v.coords)
}

// Magnitude returns the distance from the endpoint to the origin of a given
// vector.
func (v *Vector) Magnitude() float64 {
	var l float64
	for _, c := range v.coords {
		l += c * c
	}

	if l == 0 {
		return 0
	}

	return math.Sqrt(l)
}

// Normalize normalizes, i.e. divides each coodinate with its magnitude, a given
// vector.
func (v *Vector) Normalize() error {
	mag := v.Magnitude()
	if mag == 0 {
		return ErrMagZero
	}

	for k := range v.coords {
		v.coords[k] /= mag
	}

	return nil
}

// Add adds each coordinate of two given vectors.
func (v *Vector) Add(w *Vector) error {
	if len(v.coords) != len(w.coords) {
		return ErrDimsNotEqual
	}

	for k := range v.coords {
		v.coords[k] += w.coords[k]
	}

	return nil
}

// Sub substracts each coordinate of two given vectors.
func (v *Vector) Sub(w *Vector) error {
	if len(v.coords) != len(w.coords) {
		return ErrDimsNotEqual
	}

	for k := range v.coords {
		v.coords[k] -= w.coords[k]
	}

	return nil
}

// Mul multiplies each coordinate of two given vectors.
func (v *Vector) Mul(w *Vector) error {
	if len(v.coords) != len(w.coords) {
		return ErrDimsNotEqual
	}

	for k := range v.coords {
		v.coords[k] *= w.coords[k]
	}

	return nil
}

// Div divides each coordinate of two given vectors.
func (v *Vector) Div(w *Vector) error {
	if len(v.coords) != len(w.coords) {
		return ErrDimsNotEqual
	}

	for k := range v.coords {
		v.coords[k] /= w.coords[k]
	}

	return nil
}

// Dot returns the dot product (scalar product) of two given vectors.
func (v *Vector) Dot(w *Vector) (float64, error) {
	if len(v.coords) != len(w.coords) {
		return 0, ErrDimsNotEqual
	}

	var d float64
	for k := range v.coords {
		d += v.coords[k] * w.coords[k]
	}

	return d, nil
}

// Scale scales each coordinate in a given vector with a given scalar value.
func (v *Vector) Scale(scalar float64) {
	for k := range v.coords {
		v.coords[k] *= scalar
	}
}
