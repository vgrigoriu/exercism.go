// Package phonenumber solves the Phone Number problem from Exercism.
package phonenumber

import (
	"fmt"
	"regexp"
)

var areaCodeRegex = regexp.MustCompile(`^[+1 (]*([2-9]\d\d)(.*)$`)
var exchangeCodeRegex = regexp.MustCompile(`^[) .-]*([2-9]\d\d)(.*)$`)
var subscriberRegex = regexp.MustCompile(`^[ .-]*(\d{4})\s*$`)

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
	matches := areaCodeRegex.FindStringSubmatch(input)
	if len(matches) == 0 {
		return phoneNumber{}, fmt.Errorf("could not find area code")
	}
	areaCode := matches[1]
	rest := matches[2]

	matches = exchangeCodeRegex.FindStringSubmatch(rest)
	if len(matches) == 0 {
		return phoneNumber{}, fmt.Errorf("could not find exchange code")
	}
	exchangeCode := matches[1]
	rest = matches[2]

	matches = subscriberRegex.FindStringSubmatch(rest)
	if len(matches) == 0 {
		return phoneNumber{}, fmt.Errorf("could not find subscriber number")
	}
	subscriber := matches[1]

	return phoneNumber{areaCode, exchangeCode, subscriber}, nil
}

func (n *phoneNumber) number() string {
	return fmt.Sprintf("%s%s%s", n.areaCode, n.exchangeCode, n.subscriber)
}

func (n *phoneNumber) format() string {
	return fmt.Sprintf("(%s) %s-%s", n.areaCode, n.exchangeCode, n.subscriber)
}
