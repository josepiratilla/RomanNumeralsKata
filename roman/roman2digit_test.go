package roman

import "testing"

func TestCheckRoman2DigitValues(t *testing.T) {

	for expected := 1; expected < 3001; expected += 1 {
		input, err := Digit2Roman(expected)
		if err != nil {
			t.Fatal("Error at Digit2Roman when executing the Test")
		}
		actual, err := Roman2Digit(input)
		if err != nil {
			t.Errorf("Roman2Digit: Unexpected error response when converting %s\n", input)
		}
		if actual != expected {
			t.Errorf("Roman2Digit: Expected: %s -> %d\nObtained: %d\n", input, expected, actual)
		}
	}
}
