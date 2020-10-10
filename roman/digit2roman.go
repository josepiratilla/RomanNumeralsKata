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

	thousands := digit / 1000
	cents := (digit % 1000) / 100
	tens := (digit % 100) / 10
	units := digit % 10

	out := processThousands(thousands)
	out += processCents(cents)
	out += processTens(tens)
	out += processUnits(units)

	return out, nil
}

func processThousands(digit int) string {
	return processDigit(digit, "M", "-", "")
}

func processCents(digit int) string {
	return processDigit(digit, "C", "D", "M")
}

func processTens(digit int) string {
	return processDigit(digit, "X", "L", "C")
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
