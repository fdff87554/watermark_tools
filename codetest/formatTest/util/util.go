package util

import (
	"errors"
	"fmt"
)

func MatrixDot(m1, m2 [][]float64) ([][]float64, error) {

	shape1Row, shape1Col, shape2Row, shape2Col := len(m1), len(m1[0]), len(m2), len(m2[0])
	if shape1Col != shape2Row {
		return nil, errors.New(fmt.Sprintf("ValueError: shapes (%d,%d) and (%d,%d) not aligned: %d (dim 1) != %d (dim 0)", shape1Row, shape1Col, shape2Row, shape2Col, shape1Col, shape2Row))
	}
	m3 := make([][]float64, shape1Row)
	for i := range m3 {
		m3[i] = make([]float64, shape2Col)
	}

	for i := 0; i < shape1Row; i++ {
		for j := 0; j < shape2Col; j++ {
			for k := 0; k < shape1Col; k++ {
				m3[i][j] = m3[i][j] + m1[i][k]*m2[k][j]
			}
		}
	}

	return m3, nil
}

func MatrixTranfer(m1 [][]float64) [][]float64 {
	shapeRow, shapeCol := len(m1), len(m1[0])
	m2 := make([][]float64, shapeCol)
	for i := range m2 {
		m2[i] = make([]float64, shapeRow)
	}

	for i := 0; i < shapeRow; i++ {
		for j := 0; j < shapeCol; i++ {
			m2[i][j] = m2[j][i]
		}
	}

	return m2
}