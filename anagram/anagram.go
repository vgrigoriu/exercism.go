// Package anagram solves the Anagram problem from Exercism.
package anagram

import (
	"sort"
	"strings"
)

// Detect finds anagrams of subject among candidates.
func Detect(subject string, candidates []string) []string {
	lcSubject := strings.ToLower(subject)
	sortedSubject := sortRunes(lcSubject)

	result := make([]string, 0)
	for _, candidate := range candidates {
		lcCandidate := strings.ToLower(candidate)
		if lcCandidate == lcSubject {
			continue
		}
		if sortRunes(lcCandidate) != sortedSubject {
			continue
		}
		result = append(result, candidate)
	}

	return result
}

func sortRunes(input string) string {
	runes := []rune(input)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}
