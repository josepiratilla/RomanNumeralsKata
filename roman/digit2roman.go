package roman

import (
	"errors"
	"strings"
)

var romanChar = [...]struct {
	unit string
	five string
}{
	{
		"I",
		"V",
	},
	{
		"X",
		"L",
	},
	{
		"C",
		"D",
	},
	{
		"M",
		"-",
	},
	{
		"-",
		"-",
	},
}

type process struct {
	out       string
	remainder int
}

var errTooBig = errors.New("Number over 3000")

//Digit2Roman converts the input into a roman numeral string.
//It can only work with numbers up to 3000.
func Digit2Roman(digit int) (string, error) {
	if digit > 3000 {
		return "", errTooBig
	}

	p := new(process)
	p.remainder = digit
	for i := 0; p.remainder > 0; i++ {
		p.processDigit(i)
	}
	return p.out, nil
}

func (p *process) processDigit(position int) {
	digit := p.remainder % 10
	p.remainder = p.remainder / 10
	p.out = processDigit(digit, romanChar[position].unit, romanChar[position].five, romanChar[position+1].unit) + p.out
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
