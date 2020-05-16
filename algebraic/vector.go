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
)

// NewVector creates a new instance of a Vector with a given dimension and a
// slice of coordinates, and returns a pointer to it. If the dimension is
// greater than the number of coordinates, the remaining indexes of the vector
// will be zero-filled. If the dimension is less, the remaining coordinates will
// be ignored.
func NewVector(dim int, coords ...float64) (*Vector, error) {
	if dim < 0 {
		return nil, ErrInvalidDim
	}

	cs := make([]float64, dim)
	var i int
	for i < dim {
		if i < len(coords) {
			cs[i] = coords[i]
		} else {
			cs[i] = 0.0
		}
		i++
	}

	return &Vector{
		coords: cs,
	}, nil
}

// NewZeroVector creates a new instance of a zero-filled vector with a given
// dimension.
func NewZeroVector(dim int) (*Vector, error) {
	if dim < 0 {
		return nil, ErrInvalidDim
	}

	cs := make([]float64, dim)
	return &Vector{
		coords: cs,
	}, nil
}

// NewUnitVector creates a new instance of a unit vector with a given  with a
// given dimension.
func NewUnitVector(dim, coord int) (*Vector, error) {
	if dim < 0 {
		return nil, ErrInvalidDim
	}

	if dim < coord {
		return nil, ErrCoordOutOfBounds
	}

	cs := make([]float64, dim)
	cs[coord] = 1.0

	return &Vector{
		coords: cs,
	}, nil
}

// Vector defines a vector structure with a slice of floating point coordinates.
type Vector struct {
	coords []float64
}

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
func (v *Vector) Add(w *Vector) *Vector {
	rv := &Vector{
		coords: make([]float64, v.Dimension()),
	}

	for k, c := range v.coords {
		rv.coords[k] = c + w.coords[k]
	}

	return rv
}

// Sub substracts each coordinate of two given vectors.
func (v *Vector) Sub(w *Vector) *Vector {
	rv := &Vector{
		coords: make([]float64, v.Dimension()),
	}

	for k, c := range v.coords {
		rv.coords[k] = c - w.coords[k]
	}

	return rv
}

// Mul multiplies each coordinate of two given vectors.
func (v *Vector) Mul(w *Vector) *Vector {
	rv := &Vector{
		coords: make([]float64, v.Dimension()),
	}

	for k, c := range v.coords {
		rv.coords[k] = c * w.coords[k]
	}

	return rv
}

// Div divides each coordinate of two given vectors.
func (v *Vector) Div(w *Vector) *Vector {
	rv := &Vector{
		coords: make([]float64, v.Dimension()),
	}

	for k, c := range v.coords {
		rv.coords[k] = c / w.coords[k]
	}

	return rv
}

// Dot returns the dot product (scalar product) of two given vectors.
func (v *Vector) Dot(w *Vector) float64 {
	var d float64
	for k, c := range v.coords {
		d += c * w.coords[k]
	}

	return d
}

// Scale scales each coordinate in a given vector with a given scalar value.
func (v *Vector) Scale(scalar float64) *Vector {
	rv := &Vector{
		coords: make([]float64, v.Dimension()),
	}

	for k, c := range v.coords {
		rv.coords[k] = c * scalar
	}

	return rv
}

func (v *Vector) GetCoord(coord int) float64 {
	return v.coords[coord]
}
