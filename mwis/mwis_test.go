package mwis

import "testing"

func TestMWIS(t *testing.T) {
	cases := []struct {
		weights []int
		want    []int
	}{
		{
			[]int{1, 4, 5, 4},
			[]int{1, 3},
		},
		{
			[]int{4, 1, 5, 4},
			[]int{0, 2},
		},
	}

	for _, c := range cases {
		if got := MWIS(c.weights); !isEqual(got, c.want) {
			t.Errorf("MWIS(%v) = %v, want %v", c.weights, got, c.want)
		}
	}
}

func isEqual(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
