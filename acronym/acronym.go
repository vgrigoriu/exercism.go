// Package acronym solves the Acronym problem from Exercism.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns the acronym of s.
func Abbreviate(s string) string {
	prevSeparator := true
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) && prevSeparator {
			result.WriteRune(unicode.ToUpper(r))
		}
		prevSeparator = r == ' ' || r == '-'
	}

	return result.String()
}
