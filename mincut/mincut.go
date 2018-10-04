package mincut

import (
	"math"
	"math/rand"
)

// Graph represents a graph, where map[head]map[tail]edge_count
type Graph map[int]map[int]int

// Edge represents an edge between vertex a and vertex b.
type Edge struct{ a, b int }

// WithEdges first converts the slice of Edges into Graph, and call WithGraph with it.
func WithEdges(edges []Edge) int {
	g := Graph{}
	for _, e := range edges {
		if _, ok := g[e.a]; !ok {
			g[e.a] = map[int]int{}
		}
		if _, ok := g[e.b]; !ok {
			g[e.b] = map[int]int{}
		}
		g[e.a][e.b]++
		g[e.b][e.a]++
	}
	return WithGraph(g)
}

// WithGraph finds the minimum number of edges to cut a graph into two subgraphs
// using randomized contraction algorithm.
func WithGraph(g Graph) int {
	n := len(g)
	t := n * n * int(math.Log2(float64(n)))
	best := len(toEdges(g))
	for t > 0 {
		g1 := clone(g)
		cut := try(g1)
		if cut < best {
			best = cut
		}
		t--
	}
	return best
}

func toEdges(g Graph) []Edge {
	mark := map[Edge]bool{}
	edges := []Edge{}
	for h, ts := range g {
		for t, cnt := range ts {
			// NOTE: Two different entries to the same vertex are
			// considered to be distinct edges.
			if !mark[Edge{t, h}] {
				e := Edge{h, t}
				mark[e] = true
				for i := 0; i < cnt; i++ {
					edges = append(edges, e)
				}
			}
		}
	}
	return edges
}

func clone(g Graph) Graph {
	g1 := make(Graph)
	for h, ts := range g {
		ts1 := make(map[int]int)
		for t, cnt := range ts {
			ts1[t] = cnt
		}
		g1[h] = ts1
	}
	return g1
}

func try(g Graph) int {
	for len(g) > 2 {
		edges := toEdges(g)
		e := pickEdge(edges)
		merge(g, e)
	}
	return len(toEdges(g))
}

func pickEdge(edges []Edge) Edge {
	i := rand.Intn(len(edges))
	return edges[i]
}

func merge(g Graph, e Edge) {
	ats := g[e.a]
	bts := g[e.b]
	for t, cnt := range bts {
		if t == e.a {
			continue
		}
		// Migrate edges from e.b to e.a.
		ats[t] += cnt

		// Replace e.b with e.a in g[t]
		ots := g[t]
		ots[e.a] += cnt
		delete(ots, e.b)
	}
	delete(g, e.b)
	delete(ats, e.b)
}
