package karatsuba

import "testing"

func TestIsEqual(t *testing.T) {
	tests := []struct {
		a    digits
		b    digits
		want bool
	}{
		{nil, nil, true},
		{nil, digits{}, false},
		{digits{}, digits{}, true},
		{digits{}, digits{1}, false},
		{digits{1, 2, 3}, digits{1, 2, 3}, true},
	}

	for _, test := range tests {
		ab := test.a.isEqual(test.b)
		ba := test.b.isEqual(test.a)

		if !(ab == ba && ba == test.want) {
			t.Logf("%#v.isEqual(%#v) = %t", test.a, test.b, ab)
			t.Logf("%#v.isEqual(%#v) = %t", test.b, test.a, ba)
			t.Fail()
		}
	}
}

func TestToDigits(t *testing.T) {
	tests := []struct {
		input string
		want  digits
		ok    bool
	}{
		{"", digits{}, true},
		{"0", digits{0}, true},
		{"12", digits{2, 1}, true},
		{"a", nil, false},
	}

	for _, test := range tests {
		got, err := toDigits(test.input)
		ok := err == nil
		if !got.isEqual(test.want) || ok != test.ok {
			t.Errorf("toDigits(%q) = %v, %v", test.input, got, err)
		}
	}
}

func TestNumber(t *testing.T) {
	tests := []struct {
		input digits
		want  string
	}{
		{nil, ""},
		{digits{0}, "0"},
		{digits{1}, "1"},
		{digits{2, 1}, "12"},
	}

	for _, test := range tests {
		if got := test.input.number(); got != test.want {
			t.Errorf("%#v.Number() = %q", test.input, got)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want string
	}{
		{"0", "0", "0"},
		{"1", "0", "1"},
		{"1", "1", "2"},
		{"1", "9", "10"},
		{"10", "15", "25"},
		{"22", "93", "115"},
	}

	for _, test := range tests {
		a, _ := toDigits(test.a)
		b, _ := toDigits(test.b)
		ab := a.add(b).number()
		ba := b.add(a).number()
		if !(ab == ba && ba == test.want) {
			t.Logf("%#v.add(%#v) = %q", test.a, test.b, ab)
			t.Logf("%#v.add(%#v) = %q", test.b, test.a, ba)
			t.Fail()
		}
	}
}
