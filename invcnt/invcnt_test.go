package invcnt

import (
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		ar   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
		{[]int{2, 1}, 1},
		{[]int{3, 2, 1}, 3},
		{[]int{1, 3, 5, 2, 4, 6}, 3},
	}

	for _, test := range tests {
		if got := Count(test.ar); got != test.want {
			t.Errorf("Count(%v) = %d, want %d", test.ar, got, test.want)
		}
	}
}
