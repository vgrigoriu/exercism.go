// Package matrix solves the Matrix problem from Exercism.
package matrix

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Matrix is a matrix of integers.
type Matrix [][]int

// New builds a new Matrix by parsing the input string, or returns an error.
func New(input string) (Matrix, error) {
	if strings.HasSuffix(input, "\n") {
		return nil, fmt.Errorf("for some reason input cannot end in a new line")
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	cols := 0
	m := make([][]int, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		numbers := strings.Split(line, " ")

		// first line: set cols
		if cols == 0 {
			cols = len(numbers)
		}

		// other lines: check same number of elements
		if cols != len(numbers) {
			return nil, fmt.Errorf("different number of entries, %d vs. %d", len(numbers), cols)
		}

		// parse each number
		row := make([]int, 0, cols)
		for _, n := range numbers {
			val, err := strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("cannot parse %s as int: %v", n, err)
			}
			row = append(row, val)
		}
		m = append(m, row)
	}

	return Matrix(m), nil
}

// Rows returns a copy of the rows in the matrix.
func (m Matrix) Rows() [][]int {
	rows := len(m)
	cols := len(m[0])
	result := make([][]int, 0, rows)
	for _, mRow := range m {
		row := make([]int, cols)
		copy(row, mRow)
		result = append(result, row)
	}

	return result
}

// Cols returns a copy of the matrix, transposed.
func (m Matrix) Cols() [][]int {
	rows := len(m[0])
	cols := len(m)
	result := make([][]int, 0, rows)

	// create empty matrix
	for range m[0] {
		result = append(result, make([]int, cols))
	}

	// copy values transposed
	for i := range m {
		for j, val := range m[i] {
			result[j][i] = val
		}
	}

	return result
}

// Set sets the matrix entry row, col to value val.
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || len(m) <= row {
		// row out of range
		return false
	}

	if col < 0 || len(m[0]) <= col {
		// col out of range
		return false
	}

	m[row][col] = val
	return true
}
