package karatsuba

import "testing"

func TestIsEqual(t *testing.T) {
	tests := []struct {
		x    digits
		y    digits
		want bool
	}{
		{nil, nil, true},
		{nil, digits{}, false},
		{digits{}, digits{}, true},
		{digits{}, digits{1}, false},
		{digits{1, 2, 3}, digits{1, 2, 3}, true},
	}

	for _, test := range tests {
		xy := test.x.isEqual(test.y)
		yx := test.y.isEqual(test.x)

		if !(xy == yx && yx == test.want) {
			t.Logf("%#v.isEqual(%#v) = %t", test.x, test.y, xy)
			t.Logf("%#v.isEqual(%#v) = %t", test.y, test.x, yx)
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
		x, y, want string
	}{
		{"0", "0", "0"},
		{"1", "0", "1"},
		{"1", "1", "2"},
		{"1", "9", "10"},
		{"10", "15", "25"},
		{"22", "93", "115"},
	}

	for _, test := range tests {
		x, _ := toDigits(test.x)
		y, _ := toDigits(test.y)
		xy := x.add(y).number()
		yx := y.add(x).number()
		if !(xy == yx && yx == test.want) {
			t.Logf("%#v.add(%#v) = %q", test.x, test.y, xy)
			t.Logf("%#v.add(%#v) = %q", test.y, test.x, yx)
			t.Fail()
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		x, y, want string
	}{
		{"1", "0", "1"},
		{"1", "1", "0"},
		{"10", "1", "9"},
		{"100", "1", "99"},
		{"1234", "456", "778"},
		{"1000", "191", "809"},
	}

	for _, test := range tests {
		x, _ := toDigits(test.x)
		y, _ := toDigits(test.y)
		got := x.sub(y).number()
		if got != test.want {
			t.Errorf("%#v.sub(%#v) = %q", test.x, test.y, got)
		}
	}
}

func TestShift(t *testing.T) {
	tests := []struct {
		ds   digits
		k    int
		want digits
	}{
		{digits{1}, 1, digits{0, 1}},
		{digits{1, 2}, 2, digits{0, 0, 1, 2}},
		{digits{1, 2, 3}, 3, digits{0, 0, 0, 1, 2, 3}},
	}

	for _, test := range tests {
		got := test.ds.shift(test.k)
		if !got.isEqual(test.want) {
			t.Errorf("%#v.shift(%d) = %#v", test.ds, test.k, got)
		}
	}
}

func TestNoZeros(t *testing.T) {
	tests := []struct {
		ds   digits
		want digits
	}{
		{digits{1}, digits{1}},
		{digits{1, 0}, digits{1}},
		{digits{1, 2, 3, 0, 0}, digits{1, 2, 3}},
	}

	for _, test := range tests {
		got := test.ds.noZeros()
		if !got.isEqual(test.want) {
			t.Errorf("%#v.noZeros() = %#v", test.ds, got)
		}
	}

}

func TestEnsureN(t *testing.T) {
	tests := []struct {
		ds   digits
		n    int
		want digits
	}{
		{digits{1}, 1, digits{1}},
		{digits{1}, 2, digits{1, 0}},
		{digits{1, 2, 3}, 5, digits{1, 2, 3, 0, 0}},
	}

	for _, test := range tests {
		got := test.ds.ensureN(test.n)
		if !got.isEqual(test.want) {
			t.Errorf("%#v.ensureN(%d) = %#v", test.ds, test.n, got)
		}
	}

}
