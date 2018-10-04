package mincut

import (
	"testing"
)

func TestWithEdges(t *testing.T) {
	cases := []struct {
		edges []Edge
		want  int
	}{
		{[]Edge{{1, 2}}, 1},
		{[]Edge{{1, 2}, {2, 3}, {2, 4}}, 1},
		{[]Edge{{1, 2}, {1, 3}, {2, 3}, {2, 4}, {3, 4}}, 2},
		{[]Edge{{1, 2}, {2, 3}, {3, 4},
			{5, 6}, {6, 7}, {7, 8},
			{1, 6}, {2, 5}, {3, 8}, {4, 7},
		}, 2},
	}

	for _, c := range cases {
		if got := WithEdges(c.edges); got != c.want {
			t.Errorf("WithEdges(%v) = %d, want %d", c.edges, got, c.want)
		}
	}
}
