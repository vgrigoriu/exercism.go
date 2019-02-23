// Package atbash solves the Atbash Cipher problem from Exercism.
package atbash

import (
	"strings"
)

// Atbash returns the Atbash encryption of a plain text.
func Atbash(plain string) string {
	var result strings.Builder
	cipherLength := 0
	for _, r := range plain {
		ok, sr := substitute(r)
		if !ok {
			// skip non-letters and non-digits
			continue
		}
		if shouldAddSeparator(cipherLength) {
			result.WriteRune(' ')
		}
		result.WriteRune(sr)
		cipherLength++
	}

	return result.String()
}

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func substitute(r rune) (bool, rune) {
	if '0' <= r && r <= '9' {
		return true, r
	}
	var index int
	if 'a' <= r && r <= 'z' {
		index = int(r - 'a')
	} else if 'A' <= r && r <= 'Z' {
		index = int(r - 'A')
	} else {
		return false, 0
	}
	return true, rune(alphabet[len(alphabet)-index-1])
}

func shouldAddSeparator(length int) bool {
	return length > 0 && length%5 == 0
}
