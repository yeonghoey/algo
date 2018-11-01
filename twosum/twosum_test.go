package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	cases := []struct {
		a    []int
		t    int
		want int
	}{
		{[]int{}, 0, 0},
		{[]int{1, 1}, 2, 1},
		{[]int{1, 2, 1}, 3, 1},
		{[]int{1, 2, 3, 2}, 4, 2},
	}

	for _, c := range cases {
		if got := TwoSum(c.a, c.t); got != c.want {
			t.Errorf("TwoSum(%v) = %d, want %d", c.a, got, c.want)
		}
	}
}
