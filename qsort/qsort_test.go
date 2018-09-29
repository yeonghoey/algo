package qsort

import (
	"sort"
	"testing"
)

func TestFirst(t *testing.T) {
	cases := []struct {
		a     []int
		count int
	}{
		{[]int{1}, 0},
		{[]int{2, 1}, 1},
		{[]int{2, 3, 1}, 2},
		{[]int{3, 2, 1}, 3},
	}

	for _, c := range cases {
		a := make([]int, len(c.a))
		copy(a, c.a)

		count := First(a)

		if !sort.IntsAreSorted(a) {
			t.Errorf("First(%v) did not sort properly. Got %v", c.a, a)
		}
		if count != c.count {
			t.Errorf("First(%v) = %d, want %d", c.a, count, c.count)
		}
	}

}
