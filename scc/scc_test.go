package scc

import "testing"

func TestSCC(t *testing.T) {
	cases := []struct {
		edges []Edge
		want  []int
	}{
		{
			[]Edge{{1, 2}, {2, 3}},
			[]int{1, 1, 1},
		},
		{
			[]Edge{{1, 2}, {2, 1}, {2, 3}},
			[]int{2, 1},
		},
		{
			[]Edge{
				{5, 2}, {2, 8}, {8, 5}, {8, 6},
				{6, 9}, {9, 3}, {3, 6}, {9, 7},
				{7, 1}, {1, 4}, {4, 7},
			},
			[]int{3, 3, 3},
		},
	}

	for _, c := range cases {
		if got := SCC(c.edges); !isEqual(got, c.want) {
			t.Errorf("SCC(%v) = %v, want %v", c.edges, got, c.want)
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
