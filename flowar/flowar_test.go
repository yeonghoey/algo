package flowar

import "testing"

func TestFloydWarshallFrom1(t *testing.T) {
	edges := []Edge{
		{1, 2, 2},
		{1, 3, 4},
		{2, 3, 1},
		{2, 4, 2},
		{3, 5, 4},
		{4, 5, 2},
	}

	cases := []struct {
		a, b   int
		length int
	}{
		{1, 2, 2},
		{1, 3, 3},
		{1, 4, 4},
		{1, 5, 6},
	}

	shortestPaths, _ := FloydWarshall(edges)

	for _, c := range cases {
		length, ok := shortestPaths.Get(c.a, c.b)
		if !(ok && length == c.length) {
			t.Errorf("FloydWarshallFrom 1 to %d = %d(%t), want %d",
				c.b, length, ok, c.length)
		}
	}
}

func TestFloydWarshallNegCycle(t *testing.T) {
	edges := []Edge{
		{1, 2, -1},
		{2, 3, -1},
		{3, 1, -1},
	}

	_, negCycle := FloydWarshall(edges)
	if !negCycle {
		t.Errorf("Failed to detect a negative cycle")
	}
}
