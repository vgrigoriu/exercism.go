// Package strand solves the RNA Transcription problem from Exercism.
package strand

import "strings"

// ToRNA returns the RNA complement of a DNA string.
// Requires: dna to be a valid DNA string
// Guarantees: the result is the RNA complement of DNA
func ToRNA(dna string) string {
	sb := strings.Builder{}
	for _, ncl := range dna {
		sb.WriteRune(complement[ncl])
	}

	return sb.String()
}

var complement = map[rune]rune{
	'C': 'G',
	'G': 'C',
	'T': 'A',
	'A': 'U',
}
