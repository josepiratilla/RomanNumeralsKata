package roman

import (
	"errors"
)

var errNotARoman = errors.New("Not a roman numeral")

//Roman2Digit converts a roman number to an integuer number
func Roman2Digit(roman string) (int, error) {

	runes := []rune(roman)
	max := len(runes)
	out := 0
	stage := 3 //stage 3 means processing thousands, 2 processing hundreds... 0 units
	pow10 := 1000
	skip := false
	for i, r := range runes {
		if skip {
			skip = false
			continue
		}
		found := false
		for stage >= 0 {
			switch string(r) {
			case romanChar[stage].unit:
				if i < max-1 {
					switch string(runes[i+1]) {
					case romanChar[stage].five:
						found = true
						out += pow10 * 4
						skip = true
						break
					case romanChar[stage+1].unit:
						found = true
						out += pow10 * 9
						skip = true
						break
					}
				}
				if found {
					break
				}
				out += pow10
				found = true
			case romanChar[stage].five:
				out += pow10 * 5
				found = true
			}
			if found {
				break
			}
			stage--
			pow10 = pow10 / 10
		}
	}
	return out, nil
}
