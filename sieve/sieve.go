// Package sieve solves the Sieve problem from Exercism.
package sieve

// Sieve implements the Sieve of Erathostenes.
func Sieve(limit int) []int {
	isComposite := make([]bool, limit+1)
	result := make([]int, 0)
	for i := range isComposite {
		if i < 2 {
			continue
		}
		if !isComposite[i] {
			mark(isComposite, i)
			result = append(result, i)
		}
	}

	return result
}

func mark(isComposite []bool, n int) {
	for i := 2 * n; i < len(isComposite); i += n {
		isComposite[i] = true
	}
}
