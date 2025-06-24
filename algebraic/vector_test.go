package algebraic_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/madshov/data-structures/algebraic"
)

func TestNewVector(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		dim    uint
		coords []float64
		want   []float64
	}{
		"should return a new vector with dimension 3 and set values": {
			dim:    3,
			coords: []float64{1, 2, 3},
			want:   []float64{1, 2, 3},
		},
		"should return a new vector with dimension 4 and two set values": {
			dim:    4,
			coords: []float64{1, 2},
			want:   []float64{1, 2, 0, 0},
		},
		"should return a new vector with dimension 2 and only two set values": {
			dim:    2,
			coords: []float64{1, 2, 3},
			want:   []float64{1, 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := algebraic.NewVector(test.dim, test.coords...)

			assert.Len(got, int(test.dim))
			assert.EqualValues(test.want, got)
		})
	}
}

func TestNewZeroVector(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		dim  uint
		want []float64
	}{
		"should return a new zero vector with dimension 3": {
			dim:  3,
			want: []float64{0, 0, 0},
		},
		"should return a new zero vector with dimension 5": {
			dim:  5,
			want: []float64{0, 0, 0, 0, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := algebraic.NewZeroVector(test.dim)

			assert.Len(got, int(test.dim))
			assert.EqualValues(test.want, got)
		})
	}
}

func TestNewUnitVector(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		dim  uint
		el   uint
		want []float64
	}{
		"should return a new unit vector with dimension 3": {
			dim:  3,
			el:   0,
			want: []float64{1, 0, 0},
		},
		"should return a new unit vector with dimension 3 and el 2": {
			dim:  3,
			el:   2,
			want: []float64{0, 0, 1},
		},
		"should return a new unit vector with dimension 3 and el 3": {
			dim:  3,
			el:   3,
			want: []float64{0, 0, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := algebraic.NewUnitVector(test.dim, test.el)

			assert.Len(got, int(test.dim))
			assert.EqualValues(test.want, got)
		})
	}
}

func TestDimension(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		dim    uint
		coords []float64
		want   uint
	}{
		"should return a new vector with length 3": {
			dim:    3,
			coords: []float64{1, 2, 3},
			want:   3,
		},
		"should return a new vector with length 0": {
			dim:    0,
			coords: []float64{},
			want:   0,
		},
		"should return a new vector with length 5": {
			dim:    5,
			coords: []float64{1, 2, 3},
			want:   5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v := algebraic.NewVector(test.dim, test.coords...)
			got := v.Dimension()
			assert.EqualValues(test.want, got)
		})
	}
}

func TestMagnitude(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		v    algebraic.Vector
		want float64
	}{
		"should return the magnitude given a 3-dimentional vector": {
			v:    algebraic.NewVector(3, 1, 2, 3),
			want: math.Sqrt(14),
		},
		"should return 0 magnitude given an empty vector of length 0": {
			v:    algebraic.NewZeroVector(0),
			want: 0,
		},
		"should return 0 magnitude given a zero vector of length 3": {
			v:    algebraic.NewZeroVector(3),
			want: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.v.Magnitude()
			assert.InDelta(test.want, got, 0.01)
		})
	}
}

func TestNormalize(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		v       algebraic.Vector
		want    algebraic.Vector
		wantErr error
	}{
		"should return a normalized vector given a 3-dimensional vector": {
			v:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(3, 1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)),
		},
		"should return a normalize error given an empty vector of length 0": {
			v:       algebraic.NewZeroVector(0),
			wantErr: algebraic.ErrMagZero,
		},
		"should return a normalize error given a zero vector of length 3": {
			v:       algebraic.NewZeroVector(3),
			wantErr: algebraic.ErrMagZero,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.v.Normalize()
			if test.wantErr != nil {
				assert.ErrorIs(err, algebraic.ErrMagZero)
			} else {
				assert.InDeltaSlice(test.want, test.v, 0.01)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		v    algebraic.Vector
		w    algebraic.Vector
		want algebraic.Vector
	}{
		"should add 3-dimensional vector w to 3-dimensional vector v": {
			v:    algebraic.NewVector(3, 1, 2, 3),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(3, 2, 4, 6),
		},
		"should do nothing when adding 3-dimensional vector w to 0-dimensional zero vector v": {
			v:    algebraic.NewZeroVector(0),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewZeroVector(0),
		},
		"should add 3-dimensional vector w to 3-dimensional zero vector v": {
			v:    algebraic.NewZeroVector(3),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(3, 1, 2, 3),
		},
		"should add 5-dimensional vector w to 3-dimensional vector v": {
			v:    algebraic.NewVector(3, 1, 2, 3),
			w:    algebraic.NewVector(5, 1, 2, 3, 4, 5),
			want: algebraic.NewVector(3, 2, 4, 6),
		},
		"should add 3-dimensional vector w to 5-dimensional vector v": {
			v:    algebraic.NewVector(5, 1, 2, 3, 4, 5),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(5, 2, 4, 6, 4, 5),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.v.Add(test.w)
			assert.InDeltaSlice(test.want, test.v, 0.01)
		})
	}
}

func TestSub(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		v    algebraic.Vector
		w    algebraic.Vector
		want algebraic.Vector
	}{
		"should subtract 3-dimensional vector w from 3-dimensional vector v": {
			v:    algebraic.NewVector(3, 2, 3, 4),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(3, 1, 1, 1),
		},
		"should do nothing when subtracting 3-dimensional vector w from 0-dimensional zero vector v": {
			v:    algebraic.NewZeroVector(0),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewZeroVector(0),
		},
		"should subtract 3-dimensional vector w from 3-dimensional zero vector v": {
			v:    algebraic.NewZeroVector(3),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(3, -1, -2, -3),
		},
		"should subtract 5-dimensional vector w from 3-dimensional vector v": {
			v:    algebraic.NewVector(3, 2, 3, 4),
			w:    algebraic.NewVector(5, 1, 2, 3, 4, 5),
			want: algebraic.NewVector(3, 1, 1, 1),
		},
		"should subtract 3-dimensional vector w from 5-dimensional vector v": {
			v:    algebraic.NewVector(5, 2, 3, 4, 5, 6),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(5, 1, 1, 1, 5, 6),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.v.Sub(test.w)
			assert.InDeltaSlice(test.want, test.v, 0.01)
		})
	}
}

func TestMul(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		v    algebraic.Vector
		w    algebraic.Vector
		want algebraic.Vector
	}{
		"should multiply 3-dimensional vector w and 3-dimensional vector v": {
			v:    algebraic.NewVector(3, 1, 2, 3),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(3, 1, 4, 9),
		},
		"should do nothing when multiplying 3-dimensional vector w and 0-dimensional zero vector v": {
			v:    algebraic.NewZeroVector(0),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewZeroVector(0),
		},
		"should multiply 3-dimensional vector w and 3-dimensional zero vector v": {
			v:    algebraic.NewZeroVector(3),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewZeroVector(3),
		},
		"should multiply 5-dimensional vector w and 3-dimensional vector v": {
			v:    algebraic.NewVector(3, 1, 2, 3),
			w:    algebraic.NewVector(5, 1, 2, 3, 4, 5),
			want: algebraic.NewVector(3, 1, 4, 9),
		},
		"should multiply 3-dimensional vector w and 5-dimensional vector v": {
			v:    algebraic.NewVector(5, 1, 2, 3, 4, 5),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(5, 1, 4, 9, 4, 5),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.v.Mul(test.w)
			assert.InDeltaSlice(test.want, test.v, 0.01)
		})
	}
}

func TestDiv(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		v    algebraic.Vector
		w    algebraic.Vector
		want algebraic.Vector
	}{
		"should divide 3-dimensional vector w and 3-dimensional vector v": {
			v:    algebraic.NewVector(3, 1, 2, 3),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(3, 1, 1, 1),
		},
		"should do nothing when diving 3-dimensional vector w and 0-dimensional zero vector v": {
			v:    algebraic.NewZeroVector(0),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewZeroVector(0),
		},
		"should divide 3-dimensional vector w and 3-dimensional zero vector v": {
			v:    algebraic.NewZeroVector(3),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewZeroVector(3),
		},
		"should divide 5-dimensional vector w and 3-dimensional vector v": {
			v:    algebraic.NewVector(3, 1, 2, 3),
			w:    algebraic.NewVector(5, 1, 2, 3, 4, 5),
			want: algebraic.NewVector(3, 1, 1, 1),
		},
		"should divide 3-dimensional vector w and 5-dimensional vector v": {
			v:    algebraic.NewVector(5, 1, 2, 3, 4, 5),
			w:    algebraic.NewVector(3, 1, 2, 3),
			want: algebraic.NewVector(5, 1, 1, 1, 4, 5),
		},
		"should divide 5-dimensional vector w with 0's and 5-dimensional vector v": {
			v:    algebraic.NewVector(5, 1, 2, 3, 4, 5),
			w:    algebraic.NewVector(3, 1, 2, 3, 0, 0),
			want: algebraic.NewVector(5, 1, 1, 1, 4, 5),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.v.Div(test.w)
			assert.InDeltaSlice(test.want, test.v, 0.01)
		})
	}
}
