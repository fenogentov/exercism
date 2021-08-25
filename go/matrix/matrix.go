// Package matrix is a solution to lerning #16 (exercism.io)
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix stores a matrix
type Matrix [][]int

// New creates a matrix from a string
func New(s string) (Matrix, error) {
	sRows := strings.Split(s, "\n")
	res := make([][]int, len(sRows))
	var err error
	lenRow := 0
	for i, r := range sRows {
		row := strings.Split(strings.TrimSpace(r), " ")
		if len(row) == 0 {
			return nil, errors.New("empty row")
		}
		if i != 0 && lenRow != len(row) {
			return nil, errors.New("uneven rows")
		}
		lenRow = len(row)
		res[i] = make([]int, len(row))
		for j, el := range row {
			if res[i][j], err = strconv.Atoi(el); err != nil {
				return nil, err
			}
		}
	}
	return res, nil
}

// Rows reading each row left-to-right while moving top-to-bottom across the rows
func (m Matrix) Rows() [][]int {
	newM := make([][]int, len(m))
	for i := range m {
		newM[i] = make([]int, len(m[i]))
		for j := range m[i] {
			newM[i][j] = m[i][j]
		}
	}
	return newM
}

// Cols reading each column top-to-bottom while moving from left-to-right
func (m Matrix) Cols() [][]int {
	newM := make([][]int, len(m[0]))
	for i := range m[0] {
		newM[i] = make([]int, len(m))
		for j := range m {
			newM[i][j] = m[j][i]
		}
	}
	return newM
}

// Set sets an element of a matrix to a given value
func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 {
		return false
	}
	newM := *m
	if r >= len(newM) {
		return false
	}
	if c >= len(newM[0]) {
		return false
	}
	newM[r][c] = val
	return true
}
