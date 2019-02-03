// Package wordcount solves the Word Count problem from Exercism
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency represents the number of occurencies of words in a phrase.
type Frequency map[string]int

// WordCount counts the occurencies of each word in phrase.
func WordCount(phrase string) Frequency {
	result := Frequency{}
	isSeparator := func(c rune) bool {
		isInWord := unicode.IsLetter(c) || unicode.IsDigit(c) || c == '\''
		return !isInWord
	}
	words := strings.FieldsFunc(strings.ToLower(phrase), isSeparator)
	for _, word := range words {
		word = removeQuotes(word)
		result[word]++
	}
	return result
}

func removeQuotes(word string) string {
	if word[0] == '\'' && word[len(word)-1] == '\'' {
		word = word[1 : len(word)-1]
	}
	return word
}
