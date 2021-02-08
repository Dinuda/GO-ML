package goml

import "fmt"

type matrix struct {
	samples [][]float64
}

func validate(slice [][]float64) (bool, error) {
	cols := len(slice[0])

	for _, row := range slice {
		if len(row) != cols {
			return false, fmt.Errorf("Matrix dimensions did not match")
		}
	}

	return true, nil
}

func (m *matrix) multiply(matrix2 [][]float64) (*matrix, error) {
	if len(m.samples[0]) != len(matrix2) {
		return &matrix{samples: m.samples}, fmt.Errorf("Inconsistent matrix supplied")
	}

	var product [][]float64
	for k, row := range m.samples {
		for i, colVal := range row {
			product[k][i] += colVal * matrix2[i][k]
		}
	}

	return &matrix{samples: product}, nil
}

func (m *matrix) transpose() *matrix {
	r := make([][]float64, len(m.samples[0]))
	for x := range r {
		r[x] = make([]float64, len(m.samples))
	}

	for y, s := range m.samples {
		for x, e := range s {
			r[x][y] = e
		}
	}

	return &matrix{samples: r}
}

func (m *matrix) inverse() {

}

func isSquare(slice [][]float64) bool {
	cols := len(slice[0])
	rows := len(slice)

	return rows == cols
}
