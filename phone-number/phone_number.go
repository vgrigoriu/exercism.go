// Package phonenumber solves the Phone Number problem from Exercism.
package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

// Number cleans up a valid NANP number by removing punctuation and the optional country code.
func Number(input string) (number string, err error) {
	n, err := fromString(input)
	if err != nil {
		return "", err
	}

	return n.number(), nil
}

// AreaCode returns the area code of a valid NANP phone number.
func AreaCode(input string) (number string, err error) {
	n, err := fromString(input)
	if err != nil {
		return "", err
	}

	return n.areaCode, nil
}

// Format formats a valid NANP phone number.
func Format(input string) (number string, err error) {
	n, err := fromString(input)
	if err != nil {
		return "", err
	}

	return n.format(), nil
}

type phoneNumber struct {
	areaCode     string
	exchangeCode string
	subscriber   string
}

type parserState int

const (
	beforeAreaCode parserState = iota
	inAreaCode
	beforeExchangeCode
	inExchangeCode
	inSubscriber
	done
)

func fromString(input string) (*phoneNumber, error) {
	state := beforeAreaCode
	var areaCode, exchangeCode, subscriber strings.Builder
	for _, r := range input {
		// skip any punctuation
		if !unicode.IsDigit(r) {
			continue
		}
		switch state {
		case beforeAreaCode:
			// first digit must be either 1 (country code) or 2-9 (first digit of area code)
			if r == '0' {
				return nil, fmt.Errorf("unexpected first digit %q", r)
			}
			if r == '1' {
				// skip optional leading country code
				break
			}
			// we have a valid first area code digit
			areaCode.WriteRune(r)
			state = inAreaCode
			break
		case inAreaCode:
			// last two digits can be any digit
			areaCode.WriteRune(r)
			if areaCode.Len() == 3 {
				// we have enough digits for the area code
				state = beforeExchangeCode
			}
			break
		case beforeExchangeCode:
			// first digit of exchange code must be 2-9
			if r == '0' || r == '1' {
				return nil, fmt.Errorf("unexpected first digit of exchange code %q", r)
			}
			exchangeCode.WriteRune(r)
			state = inExchangeCode
			break
		case inExchangeCode:
			// last two digit can be anything
			exchangeCode.WriteRune(r)
			if exchangeCode.Len() == 3 {
				// we have enough digits for the exchange code
				state = inSubscriber
			}
			break
		case inSubscriber:
			// any four digits
			subscriber.WriteRune(r)
			if subscriber.Len() == 4 {
				state = done
			}
			break
		case done:
			// already got all the digits we need
			return nil, fmt.Errorf("too many digits")
		}
	}
	if state != done {
		return nil, fmt.Errorf("not enough digits")
	}
	return &phoneNumber{areaCode.String(), exchangeCode.String(), subscriber.String()}, nil
}

func (n *phoneNumber) number() string {
	return fmt.Sprintf("%s%s%s", n.areaCode, n.exchangeCode, n.subscriber)
}

func (n *phoneNumber) format() string {
	return fmt.Sprintf("(%s) %s-%s", n.areaCode, n.exchangeCode, n.subscriber)
}
