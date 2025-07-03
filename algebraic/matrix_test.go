package algebraic_test

import (
	"testing"

	"github.com/madshov/data-structures/algebraic"
	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		rows, cols uint
		coords     []float64
		want       algebraic.Matrix
	}{
		"should return a new vector with dimension 3 and set values": {
			rows: 3,
			cols: 3,
			coords: []float64{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			want: algebraic.NewMatrix(3, 3,
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			),
		},
		"should return a new vector with dimension 3 and set and zero values": {
			rows: 3,
			cols: 3,
			coords: []float64{
				1, 2, 3,
				4, 5, 6,
				7},
			want: algebraic.NewMatrix(3, 3,
				1, 2, 3,
				4, 5, 6,
				7, 0, 0,
			),
		},
		"should return a new vector with dimension 3 and set and zero values2": {
			rows: 3,
			cols: 3,
			coords: []float64{
				1, 2, 3,
				4, 5, 6,
			},
			want: algebraic.NewMatrix(3, 3,
				1, 2, 3,
				4, 5, 6,
				0, 0, 0,
			),
		},
		"should return a new vector with dimension 3 and set values3": {
			rows: 2,
			cols: 2,
			coords: []float64{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			want: algebraic.NewMatrix(2, 2,
				1, 2,
				3, 4,
			),
		},
		"should return a new zero matrix": {
			rows: 0,
			cols: 0,
			want: algebraic.NewMatrix(0, 0),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := algebraic.NewMatrix(test.rows, test.cols, test.coords...)
			// assert.Len(got, len(test.want))
			// assert.Len(got[0], len(test.want[0]))
			assert.EqualValues(test.want, got)
		})
	}
}

func TestNewIdentityMatrix(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		rows, cols uint
		want       algebraic.Matrix
	}{
		"should return a new vector with dimension 3": {
			rows: 3,
			cols: 3,
			want: algebraic.NewMatrix(3, 3,
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
			),
		},
		"should return a new vector with dimension 3 and set values": {
			rows: 3,
			cols: 4,
			want: algebraic.NewMatrix(3, 3,
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
			),
		},
		"should return a new vector with dimension 3 and set values2": {
			rows: 4,
			cols: 3,
			want: algebraic.NewMatrix(3, 3,
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
			),
		},
		"should return a new zero matrix": {
			rows: 0,
			cols: 0,
			want: algebraic.NewMatrix(0, 0),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := algebraic.NewIdentityMatrix(test.rows, test.cols)
			// assert.Len(got, len(test.want))
			// assert.Len(got[0], len(test.want[0]))
			assert.EqualValues(test.want, got)
		})
	}
}

func TestTranspose(t *testing.T) {
	assert := assert.New(t)
	tests := map[string]struct {
		m    algebraic.Matrix
		want algebraic.Matrix
	}{
		"should return a transposed 3x2-matrix from a 2x3-matrix": {
			m: algebraic.NewMatrix(2, 3,
				1, 2, 3,
				4, 5, 6,
			),
			want: algebraic.NewMatrix(3, 2,
				1, 4,
				2, 5,
				3, 6,
			),
		},
		"should return a transposed 3x3-matrix from a 3x3-matrix": {
			m: algebraic.NewMatrix(3, 3,
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			),
			want: algebraic.NewMatrix(3, 3,
				1, 4, 7,
				2, 5, 8,
				3, 6, 9,
			),
		},
		"should return a new zero matrix": {
			m:    algebraic.NewMatrix(0, 0),
			want: algebraic.NewMatrix(0, 0),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.m.Transpose()
			assert.EqualValues(test.want, got)
		})
	}

}
