package karatsuba

import "testing"

func TestToDigits(t *testing.T) {
	tests := []struct {
		input string
		want  digits
		err   error
	}{
		{"0", []int8{0}, nil},
	}

	for _, test := range tests {
		got, err := toDigits(test.input)
		if !isEqual(got, test.want) || err != test.err {
			t.Errorf("toDigits(%q) = %v, %v", test.input, got, err)
		}
	}
}
