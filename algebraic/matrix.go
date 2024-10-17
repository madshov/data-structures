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
func NewMatrix(cols int, coords ...float64) (*Matrix, error) {
	if cols < 0 {
		return nil, ErrInvalidCols
	}

	if cols > len(coords) {
		return nil, ErrColsOutOfBounds
	}

	var (
		i, j int
		vs   []*Vector
	)

	// loop through coords in chunks of cols
	for i < len(coords) {
		j = i + cols
		// cap upper bound if it becomes too big
		if j > len(coords) {
			j = len(coords)
		}

		v, err := NewVector(cols, coords[i:j]...)
		if err != nil {
			return nil, err
		}
		vs = append(vs, v)
		i += cols
	}

	return &Matrix{
		coords: vs,
		rows:   len(vs),
		cols:   cols,
	}, nil
}

func NewIdentityMatrix(rows, cols int) (*Matrix, error) {
	if rows != cols {
		return nil, ErrNotSquare
	}

	var (
		i  int
		vs []*Vector
	)

	for i < rows {
		v, err := NewUnitVector(cols, i)
		if err != nil {
			return nil, err
		}
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
	coords     []*Vector
	rows, cols int
}

// Transpose creates and returns a new matrix with rows and columns transposed.
// |1.0  2.0  3.0|          |1.0  4.0|
// |4.0  5.0  6.0|    =>    |2.0  5.0|
//
//	|3.0  6.0|
func (m *Matrix) Transpose() *Matrix {
	var (
		coords []float64
		vs     []*Vector
		rows   = m.cols
		cols   = m.rows
	)

	var i int
	for i < rows {
		var j int
		for j < cols {
			coords = append(coords, m.coords[j].coords[i])
			j++
		}
		i++
	}

	var k int
	for k < len(coords) {
		v, _ := NewVector(cols, coords[k:k+cols]...)
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
		for _, c := range r.coords {
			fmt.Printf("%f ", c)
		}
		fmt.Println()
	}
}
