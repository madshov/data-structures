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
		dim    uint
		coords []float64
		want   float64
	}{
		"should return correct magnitude for a vector with length 3": {
			dim:    3,
			coords: []float64{1, 2, 3},
			want:   math.Sqrt(14),
		},
		"should return correct magnitude for a vector with length 0": {
			dim:    0,
			coords: []float64{},
			want:   0,
		},
		"should return correct magnitude for a zerp vector with length 3": {
			dim:    3,
			coords: []float64{0, 0, 0},
			want:   0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v := algebraic.NewVector(test.dim, test.coords...)
			got := v.Magnitude()
			assert.InDelta(test.want, got, 0.01)
		})
	}
}
