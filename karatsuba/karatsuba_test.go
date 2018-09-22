package karatsuba

import "testing"

func TestKaratsuba(t *testing.T) {
	tests := []struct {
		xs, ys, want string
	}{
		{"0", "0", "0"},
		{"1", "0", "0"},
		{"1", "1", "1"},
		{"1", "2", "2"},
		{"2", "2", "4"},
		{"9", "9", "81"},
		{"12", "34", "408"},
		{"1234", "2", "2468"},
		{"1234", "10", "12340"},
		{"2", "5678", "11356"},
		{"11", "123", "1353"},
		{"1234", "5678", "7006652"},
	}

	for _, test := range tests {
		got := Karatsuba(test.xs, test.ys)
		if got != test.want {
			t.Errorf("Karatsuba(%q, %q) = %q", test.xs, test.ys, got)
		}
	}
}
