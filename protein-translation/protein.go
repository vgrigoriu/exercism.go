// Package protein solves the Protein Translation problem from Exercism.
package protein

import (
	"errors"
)

// ErrStop is returned when you try to decode a stop codon.
var ErrStop = errors.New("stop codon")

// ErrInvalidBase is returned when you try to decode an invalid base.
var ErrInvalidBase = errors.New("invalid base")

// FromCodon translates a codon into a protein or returns an error.
func FromCodon(codon string) (string, error) {
	if protein, ok := bases[codon]; ok {
		return protein, nil
	}
	if stopCodons[codon] {
		return "", ErrStop
	}
	return "", ErrInvalidBase
}

// FromRNA translates a string of RNA into a list of proteins.
func FromRNA(rna string) ([]string, error) {
	result := []string{}
	for i := 0; i+3 <= len(rna); i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err == ErrStop {
			return result, nil
		}
		if err != nil {
			return result, err
		}
		result = append(result, protein)
	}
	return result, nil
}

var bases = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCG": "Serine",
	"UCU": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGG": "Tryptophan",
}

var stopCodons = map[string]bool{"UAA": true, "UAG": true, "UGA": true}
