// Package cryptosquare solves the Crypto Square problem from Exercism.
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode returns the square code encoding of the input.
func Encode(input string) string {
	normalized := normalize(input)
	rmax, cmax := dimensions(len(normalized))
	lines := make([][]rune, cmax)
	for line := range lines {
		lines[line] = make([]rune, rmax)
	}
	for r := 0; r < rmax; r++ {
		for c := 0; c < cmax; c++ {
			lines[c][r] = get(normalized, c, r, cmax)
		}
	}
	slines := make([]string, cmax)
	for r := range slines {
		slines[r] = string(lines[r])
	}
	return strings.Join(slines, " ")
}

func get(input []rune, r, c, cmax int) rune {
	i := c*cmax + r
	if i >= len(input) {
		return ' '
	}
	return input[i]
}

// Normalize removes spaces and punctuation from the input string
// and downcases all the letters.
func normalize(input string) []rune {
	var result []rune
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, unicode.ToLower(r))
		}
	}

	return result
}

// Dimensions return r and c so that:
// - 0 <= c - r <= 1
// - c * r >= n
// - there is no smaller c, r
func dimensions(n int) (r, c int) {
	r = int(math.Sqrt(float64(n)))
	c = r
	for c*r < n {
		if r < c {
			r++
		} else {
			c++
		}
	}
	return
}
