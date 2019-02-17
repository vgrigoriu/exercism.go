package romannumerals

import "fmt"

// ToRomanNumeral returns the roman representation of a number between 1 and 3000, or false.
func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || 3000 < arabic {
		return "", fmt.Errorf("can only convert numbers between 1 and 3000")
	}

	// units, tens, hundreds, thousands
	thousands := convertThousands(arabic / 1000)
	hundreds := convertHundreds((arabic % 1000) / 100)
	tens := convertTens((arabic % 100 / 10))
	units := convertUnits(arabic % 10)

	return fmt.Sprintf("%s%s%s%s", thousands, hundreds, tens, units), nil
}

// Requires: 0 <= thousands <= 3
func convertThousands(thousands int) string {
	return []string{"", "M", "MM", "MMM"}[thousands]
}

// Requires: 0 <= hundreds <= 9
func convertHundreds(hundreds int) string {
	return []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}[hundreds]
}

// Requires: 0 <= tens <= 9
func convertTens(tens int) string {
	return []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}[tens]
}

// Requires: 0 <= units <= 9
func convertUnits(units int) string {
	return []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}[units]
}
