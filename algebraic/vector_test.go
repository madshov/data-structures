package algebraic_test

import (
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
			v := algebraic.NewVector(test.dim, test.coords...)

			assert.Len(v, int(test.dim))
			assert.EqualValues(test.want, v)
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
			v := algebraic.NewZeroVector(test.dim)

			assert.Len(v, int(test.dim))
			assert.EqualValues(test.want, v)
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
			v := algebraic.NewUnitVector(test.dim, test.el)

			assert.Len(v, int(test.dim))
			assert.EqualValues(test.want, v)
		})
	}
}
