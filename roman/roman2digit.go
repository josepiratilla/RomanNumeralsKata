package roman

import "errors"

var errNotARoman = errors.New("Not a roman numeral")

//Roman2Digit converts a roman number to an integuer number
func Roman2Digit(roman string) (int, error) {

	runes := []rune(roman)
	out := 0
	for _, r := range runes {
		if r == 'M' {
			out += 1000
		}
	}
	return out, nil
}
