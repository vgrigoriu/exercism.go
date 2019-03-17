package prime

// Factors computes the prime factors of n.
func Factors(n int64) []int64 {
	factors := make([]int64, 0)
	var factor int64
	for factor = 2; n > 1; factor++ {
		for n%factor == 0 {
			n /= factor
			factors = append(factors, factor)
		}
	}

	return factors
}
