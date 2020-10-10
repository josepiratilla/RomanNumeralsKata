package roman

import "testing"

func TestCheckRoman2DigitValues(t *testing.T) {

	for expected := 1; expected < 3001; expected++ {
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

func TestNotRomanNumberError(t *testing.T) {
	vs := []string{
		"MIM",
		"Hola",
		"P",
		"MMMM",
		"IVIV",
		"IIV",
		"IVIX",
		"IXI",
		"CCM",
	}
	for _, v := range vs {
		_, err := Roman2Digit(v)
		if err != errNotARoman {
			t.Errorf("Roman2Digit: Invalid roman \"%s\" didn't throw the correct error.", v)
		}
	}
}
