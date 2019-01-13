// Package strain solves the Strain problem from Exercism.
package strain

// Ints is a slice of ints.
type Ints []int

// Lists is a slice of int slices.
type Lists [][]int

// Strings is a slice of strings.
type Strings []string

// Keep filters is to the ones that satisfy the predicate.
func (is Ints) Keep(p func(int) bool) Ints {
	var result Ints
	for _, i := range is {
		if p(i) {
			result = append(result, i)
		}
	}
	return result
}

// Discard filters is to the ones that don't satisfy the predicate.
func (is Ints) Discard(p func(int) bool) Ints {
	return is.Keep(func(i int) bool { return !p(i) })
}

// Keep filters ls to the ones that satisfy the predicate.
func (ls Lists) Keep(p func([]int) bool) Lists {
	var result Lists
	for _, l := range ls {
		if p(l) {
			result = append(result, l)
		}
	}
	return result
}

// Keep filters ss to the ones that satisfy the predicate.
func (ss Strings) Keep(p func(string) bool) Strings {
	var result Strings
	for _, s := range ss {
		if p(s) {
			result = append(result, s)
		}
	}
	return result
}
