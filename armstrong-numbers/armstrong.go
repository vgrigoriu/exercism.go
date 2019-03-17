package armstrong

import (
	"math"
)

// IsNumber returns true iff n is an Armstrong number.
func IsNumber(n int) bool {
	noDigits := int(math.Log10(float64(n)) + 1)
	var sum float64
	for nPrime := n; nPrime > 0; {
		digit := nPrime % 10
		nPrime /= 10
		sum += math.Pow(float64(digit), float64(noDigits))
	}

	return n == int(sum)
}
