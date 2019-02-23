// Package atbash solves the Atbash Cipher problem from Exercism.
package atbash

import (
	"strings"
)

// Atbash returns the Atbash encryption of a plain text.
func Atbash(plain string) string {
	var result strings.Builder
	for _, r := range plain {
		ok, sr := substitute(r)
		if !ok {
			// skip non-letters and non-digits
			continue
		}
		if result.Len()%6 == 5 {
			result.WriteRune(' ')
		}
		result.WriteRune(sr)
	}

	return result.String()
}

func substitute(r rune) (bool, rune) {
	if '0' <= r && r <= '9' {
		return true, r
	}
	if 'a' <= r && r <= 'z' {
		return true, 'a' + 'z' - r
	}
	if 'A' <= r && r <= 'Z' {
		return true, 'a' + 'Z' - r
	}
	return false, 0
}

func shouldAddSeparator(length int) bool {
	return length > 0 && length%5 == 0
}
