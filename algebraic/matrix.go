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

// Maxtrix defines a maxtrix structure with a slice of Vectors.
type Matrix []Vector

// NewMatrix creates a new instance of a Matrix with a given number of columns
// and a slice of coordinates, and returns a pointer to it. The function will
// fill up each row of the matrix as long as there are more coordinates. If
// there is not enough coordinates to fill the last row, the remaining will be
// zero-filled.
func NewMatrix(rows, cols uint, coords ...float64) Matrix {
	var (
		i, j uint
		dim  = uint(len(coords))
		m    Matrix
	)

	for ; i < rows; i++ {
		var cs = make([]float64, cols)
		k := (i + 1) * cols

		if j+k < dim {
			cs = coords[j : j+k]
		} else {
			if j < dim {
				cs = coords[j:]
			}
		}

		v := NewVector(cols, cs...)
		m = append(m, v)

		j += cols
	}

	return m
}

func NewIdentityMatrix(rows, cols uint) Matrix {
	// Make sure it's a square matrix, i.e. set rows and cols to which ever is
	// minimum.
	if rows < cols {
		cols = rows
	} else {
		rows = cols
	}

	var (
		i uint
		m Matrix
	)

	for ; i < rows; i++ {
		v := NewUnitVector(cols, i)
		m = append(m, v)
	}

	return m
}

// Transpose creates and returns a new matrix with rows and columns transposed.
// |1.0  2.0  3.0|          |1.0  4.0|
// |4.0  5.0  6.0|    =>    |2.0  5.0|
//
//	|3.0  6.0|
// func (m *Matrix) Transpose() *Matrix {
// 	var (
// 		coords []float64
// 		vs     []Vector
// 		rows   = m.cols
// 		cols   = m.rows
// 	)

// 	var i uint
// 	for i < rows {
// 		var j uint
// 		for j < cols {
// 			coords = append(coords, m.coords[j][i])
// 			j++
// 		}
// 		i++
// 	}

// 	var k uint
// 	for k < uint(len(coords)) {
// 		v := NewVector(cols, coords[k:k+cols]...)
// 		vs = append(vs, v)
// 		k += cols
// 	}

// 	return &Matrix{
// 		coords: vs,
// 		rows:   rows,
// 		cols:   cols,
// 	}
// }

// func (m *Matrix) Determinant																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																					() float64 {
// }

// func (m *Matrix) Inverse() error {
// 	if m.rows != m.cols {
// 		return ErrNotSquare
// 	}
// }

func (m Matrix) Print() {
	for _, r := range m {
		for _, c := range r {
			fmt.Printf("%f ", c)
		}
		fmt.Println()
	}
}
