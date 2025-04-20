package algebraic

import (
	"errors"
	"fmt"
)

// Various errors a matrix function can return.
var (
	ErrInvalidCols     = errors.New("matrix columns cannot be negative")
	ErrColsOutOfBounds = errors.New("number of coordinates are less than colums")
	ErrNotSquare       = errors.New("matrix is not square")
)

// NewMatrix creates a new instance of a Matrix with a given number of columns
// and a slice of coordinates, and returns a pointer to it. The function will
// fill up each row of the matrix as long as there are more coordinates. If
// there is not enough coordinates to fill the last row, the remaining will be
// zero-filled.
func NewMatrix(cols uint, coords ...float64) (*Matrix, error) {
	var (
		dim = uint(len(coords))
	)

	if cols > dim {
		return nil, ErrColsOutOfBounds
	}

	var (
		i, j uint
		vs   []Vector
	)

	// loop through coords in chunks of cols
	for i < dim {
		j = i + cols
		// cap upper bound if it becomes too big
		if j > dim {
			j = dim
		}

		v := NewVector(cols, coords[i:j]...)
		vs = append(vs, v)
		i += cols
	}

	return &Matrix{
		coords: vs,
		rows:   uint(len(vs)),
		cols:   cols,
	}, nil
}

func NewIdentityMatrix(rows, cols uint) (*Matrix, error) {
	if rows != cols {
		return nil, ErrNotSquare
	}

	var (
		i  uint
		vs []Vector
	)

	for i < rows {
		v := NewUnitVector(cols, i)
		vs = append(vs, v)
		i++
	}

	return &Matrix{
		coords: vs,
		rows:   rows,
		cols:   cols,
	}, nil
}

// Matrix defines a matrix structure as a slice of vectors and a row and column
// count.
type Matrix struct {
	coords     []Vector
	rows, cols uint
}

// Transpose creates and returns a new matrix with rows and columns transposed.
// |1.0  2.0  3.0|          |1.0  4.0|
// |4.0  5.0  6.0|    =>    |2.0  5.0|
//
//	|3.0  6.0|
func (m *Matrix) Transpose() *Matrix {
	var (
		coords []float64
		vs     []Vector
		rows   = m.cols
		cols   = m.rows
	)

	var i uint
	for i < rows {
		var j uint
		for j < cols {
			coords = append(coords, m.coords[j][i])
			j++
		}
		i++
	}

	var k uint
	for k < uint(len(coords)) {
		v := NewVector(cols, coords[k:k+cols]...)
		vs = append(vs, v)
		k += cols
	}

	return &Matrix{
		coords: vs,
		rows:   rows,
		cols:   cols,
	}
}

// func (m *Matrix) Determinant																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																					() float64 {
// }

// func (m *Matrix) Inverse() error {
// 	if m.rows != m.cols {
// 		return ErrNotSquare
// 	}
// }

func (m *Matrix) Print() {
	for _, r := range m.coords {
		for _, c := range r {
			fmt.Printf("%f ", c)
		}
		fmt.Println()
	}
}
