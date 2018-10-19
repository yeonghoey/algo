package dijkstra

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	cases := []struct {
		edges []Edge
		start Vertex
		want  map[Vertex]Dist
	}{
		{
			[]Edge{{1, 2, 1}, {1, 3, 4}, {2, 3, 2}, {2, 4, 6}, {3, 4, 3}},
			1,
			map[Vertex]Dist{1: 0, 2: 1, 3: 3, 4: 6},
		},
	}

	for _, c := range cases {
		got := Dijkstra(c.edges, c.start)
		for v, d := range c.want {
			if got[v] != d {
				t.Errorf("Dijkstra(%v.., (%d edges))[%d] = %d, want %d",
					c.edges[:3], len(c.edges), v, got[v], d)
			}
		}
	}
}
