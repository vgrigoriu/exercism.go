// Package phonenumber solves the Phone Number problem from Exercism.
package phonenumber

import (
	"fmt"
	"regexp"
)

// phoneNumber represents a NANP phone number.
type phoneNumber struct {
	areaCode     string
	exchangeCode string
	subscriber   string
}

// Number cleans up a valid NANP number by removing punctuation and the optional country code.
func Number(input string) (number string, err error) {
	n, err := parse(input)
	if err != nil {
		return "", err
	}

	return n.number(), nil
}

// AreaCode returns the area code of a valid NANP phone number.
func AreaCode(input string) (number string, err error) {
	n, err := parse(input)
	if err != nil {
		return "", err
	}

	return n.areaCode, nil
}

// Format formats a valid NANP phone number.
func Format(input string) (number string, err error) {
	n, err := parse(input)
	if err != nil {
		return "", err
	}

	return n.format(), nil
}

func parse(input string) (phoneNumber, error) {
	//                            (223) 456-7890
	re := regexp.MustCompile(`^[+1 (]*([2-9]\d\d)[) .-]*([2-9]\d\d)[ .-]*(\d{4})\s*$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) == 0 {
		return phoneNumber{}, fmt.Errorf("invalid phone number")
	}

	return phoneNumber{matches[1], matches[2], matches[3]}, nil
}

func (n *phoneNumber) number() string {
	return fmt.Sprintf("%s%s%s", n.areaCode, n.exchangeCode, n.subscriber)
}

func (n *phoneNumber) format() string {
	return fmt.Sprintf("(%s) %s-%s", n.areaCode, n.exchangeCode, n.subscriber)
}
