package roman

import "errors"

var errNotARoman = errors.New("Not a roman numeral")

//Roman2Digit converts a roman number to an integuer number
func Roman2Digit(roman string) (int, error) {

	runes := []rune(roman)
	max := len(runes)
	out := 0
	stage := 4 //stage 4 means processing thousands, 3 processing hundreds...
	skip := false
	for i, r := range runes {
		if skip {
			skip = false
			continue
		}
		if stage == 4 {
			if r == 'M' {
				out += 1000
				continue
			}
			stage--
		}
		if stage == 3 {
			switch r {
			case 'C':
				if i < max-1 {
					switch runes[i+1] {
					case 'D':
						out += 400
						skip = true
						continue
					case 'M':
						out += 900
						skip = true
						continue
					}
				}
				out += 100
				continue
			case 'D':
				out += 500
				continue
			}

		}
	}
	return out, nil
}
