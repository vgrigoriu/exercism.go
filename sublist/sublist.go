// Package sublist solves the Sublist problem from Exercism.
package sublist

// Relation is one of: equal, unequal, sublist, superlist
type Relation string

// Sublist compares two int slices and returns:
// - equal if they are equal
// - sublist if b contains a
// - superlist if a contains b
// - unequal otherwise
func Sublist(a, b []int) Relation {
	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				return "unequal"
			}
		}
		return "equal"
	}

	if len(a) < len(b) {
		return sublist(a, b, "sublist")
	}

	return sublist(b, a, "superlist")
}

func sublist(small, big []int, matchResult Relation) Relation {
out:
	for i := 0; i <= len(big)-len(small); i++ {
		for j := range small {
			if big[i+j] != small[j] {
				// try next i
				continue out
			}
		}
		return matchResult
	}
	return "unequal"
}
