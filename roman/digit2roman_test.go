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

func TestCheckDigit2RomanValues(t *testing.T) {
	vs := []struct {
		input    int
		expected string
	}{
		{
			1,
			"I",
		},
		{
			3,
			"III",
		},
		{
			4,
			"IV",
		},
		{
			5,
			"V",
		},
		{
			8,
			"VIII",
		},
		{
			9,
			"IX",
		},
		// {
		// 	10,
		// 	"X",
		// },
		// {
		// 	13,
		// 	"XIII",
		// },
		// {
		// 	25,
		// 	"XXV",
		// },
		// {
		// 	39,
		// 	"XXXIX",
		// },
		// {
		// 	44,
		// 	"XLIV",
		// },
		// {
		// 	68,
		// 	"LXVIII",
		// },
		// {
		// 	99,
		// 	"XCIX",
		// },
	}
	for _, v := range vs {
		actual, _ := Digit2Roman(v.input)
		if actual != v.expected {
			t.Errorf("Digit2Roman:\nOperation: %d -> \"%s\"\nExpected: \"%s\"",
				v.input, actual, v.expected)
		}
	}
}
