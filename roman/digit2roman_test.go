package roman

import "testing"

func TestNumberTooBigError(t *testing.T) {
	vError := []int{3001, 10000}
	vOK := []int{0, 1034, 3000}
	for _, v := range vError {
		_, err := Digit2Roman(v)
		if err != errTooBig {
			t.Errorf("Digit2Roman didn't throw an error for input %d", v)
		}
	}
	for _, v := range vOK {
		_, err := Digit2Roman(v)
		if err != nil {
			t.Errorf("Digit2Roman threw an error for input %d", v)
		}
	}

}
