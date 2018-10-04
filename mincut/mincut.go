package mincut

import (
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

	t := n * n
	ch := make(chan int, t)
	for i := 0; i < t; i++ {
		go func() {
			g1 := clone(g)
			cut := try(g1)
			ch <- cut
		}()
	}
	best := len(toEdges(g))
	for i := 0; i < t; i++ {
		cut := <-ch
		if cut < best {
			best = cut
		}
	}
	return best
}

func try(g Graph) int {
	edges := toEdges(g)
	for len(g) > 2 {
		e := pickEdge(edges)
		g, edges = merge(g, edges, e)
	}
	return len(edges)
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

func pickEdge(edges []Edge) Edge {
	i := rand.Intn(len(edges))
	return edges[i]
}

func merge(g Graph, edges []Edge, e Edge) (Graph, []Edge) {
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

	for i := 0; i < len(edges); {
		e1 := edges[i]
		if (e1.a == e.a && e1.b == e.b) || (e1.b == e.a && e1.a == e.b) {
			edges[i] = edges[len(edges)-1]
			edges = edges[:len(edges)-1]
			continue
		}
		if edges[i].a == e.b {
			edges[i].a = e.a
		}
		if edges[i].b == e.b {
			edges[i].b = e.a
		}
		i++
	}

	return g, edges
}
