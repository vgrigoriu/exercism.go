// Package lsproduct solves the Largest Series Product problem from Exercism.
package lsproduct

import (
	"fmt"
	"unicode"
)

// LargestSeriesProduct computes the largest product for a contiguous substring of digits of length n
func LargestSeriesProduct(digits string, n int) (int64, error) {
	var result int64
	if n < 0 {
		return 0, fmt.Errorf("n must be greater than 0, but got %d", n)
	}
	if n > len(digits) {
		return 0, fmt.Errorf("n is bigger than no. of digits: %d vs. %q", n, digits)
	}
	if n == 0 {
		return 1, nil
	}
	for i := range digits[:len(digits)-n+1] {
		current := int64(1)
		for _, r := range digits[i : i+n] {
			if !unicode.IsDigit(r) {
				return 0, fmt.Errorf("unexpected input: %q", r)
			}
			digit := int64(r - '0')
			if digit == 0 {
				current = 0
				break
			}
			current *= digit
		}
		if current > result {
			result = current
		}
	}
	return result, nil
}
