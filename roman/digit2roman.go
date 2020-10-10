package roman

import (
	"errors"
	"strings"
)

var errTooBig = errors.New("Number over 3000")

//Digit2Roman converts the input into a roman numeral string.
//It can only work with numbers up to 3000.
func Digit2Roman(digit int) (string, error) {
	if digit > 3000 {
		return "", errTooBig
	}

	var out string

	out = out + processUnits(digit)
	return out, nil
}

func processUnits(digit int) string {

	return processDigit(digit, "I", "V", "X")
}

func processDigit(digit int, unit string, five string, ten string) string {
	if digit == 9 {
		return unit + ten
	}
	if digit == 4 {
		return unit + five
	}
	out := ""
	if digit >= 5 {

		out = out + five
		digit -= 5
	}
	out = out + strings.Repeat(unit, digit)
	return out
}
