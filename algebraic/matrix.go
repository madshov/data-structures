package algebraic

import (
	"errors"
	"fmt"
)

// Various errors a matrix function can return.
var (
	ErrInvalidCols     = errors.New("matrix columns cannot be negative")
	ErrColsOutOfBounds = errors.New("matrix coordinates are less than colums")
)

func NewMatrix(cols int, coords ...float64) (*Matrix, error) {
	if cols < 0 {
		return nil, ErrInvalidCols
	}

	if cols > len(coords) {
		return nil, ErrColsOutOfBounds
	}

	var (
		i  int
		vs []*Vector
	)

	for i+cols-1 < len(coords) {
		v := NewVector(coords[i : i+cols]...)
		vs = append(vs, v)
		i += cols
	}

	return &Matrix{
		coords: vs,
		rows:   len(vs),
		cols:   cols,
	}, nil
}

type Matrix struct {
	coords     []*Vector
	rows, cols int
}

func (m *Matrix) Transpose() (*Matrix, error) {
	var (
		i      int
		coords []float64
	)

	for i < m.cols {
		var j int
		for j < m.rows {
			coords = append(coords, m.coords[j].coords[i])
			j++
		}
		i++
	}

	n, err := NewMatrix(m.rows, coords...)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (m *Matrix) Print() {
	for _, r := range m.coords {
		for _, c := range r.coords {
			fmt.Printf("%f ", c)
		}
		fmt.Println()
	}
}
