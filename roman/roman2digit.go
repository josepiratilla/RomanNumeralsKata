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
	countUnits := 0
	foundFive := false
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
						if countUnits > 0 || foundFive {
							return 0, errNotARoman
						}
						foundFive = true
						found = true
						out += pow10 * 4
						skip = true
						break
					case romanChar[stage+1].unit:
						if countUnits > 0 || foundFive {
							return 0, errNotARoman
						}
						countUnits += 9
						found = true
						out += pow10 * 9
						skip = true
						break
					}
				}
				if found {
					break
				}
				if countUnits > 2 {
					return 0, errNotARoman
				}
				countUnits++
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
			countUnits = 0
			foundFive = false
		}
		if !found {
			return 0, errNotARoman
		}
	}
	return out, nil
}
