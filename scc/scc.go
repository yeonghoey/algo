package scc

import (
	"sort"
)

// Edge is a directed one, which represents T -> H.
type Edge struct {
	T int
	H int
}

// SCC computes strongly connected components and returns sizes of the SCCs in decreasing order.
func SCC(edges []Edge) []int {
	leaderMap := SCC2(edges)
	sccs := make(map[int]int)
	for _, v := range leaderMap {
		sccs[v]++
	}

	sizes := make([]int, 0, len(sccs))
	for _, v := range sccs {
		sizes = append(sizes, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes
}

// SCC2 computes strongly connected components and
// returns a map which associates each node to its leader node.
func SCC2(edges []Edge) map[int]int {
	g := make(map[int][]int)
	gRev := make(map[int][]int)
	for _, e := range edges {
		g[e.T] = append(g[e.T], e.H)
		gRev[e.H] = append(gRev[e.H], e.T)
	}

	// orderMap contains node:order
	orderMap := make(map[int]int)
	for _, e := range edges {
		orderMap[e.H] = e.H
		orderMap[e.T] = e.T
	}

	// leaderMap contains node:leader
	leaderMap := make(map[int]int)

	nodes := make([]int, 0, len(orderMap))
	for k := range orderMap {
		nodes = append(nodes, k)
	}

	dfsLoop := func(g map[int][]int) {
		sort.Slice(nodes, func(i, j int) bool {
			return orderMap[nodes[i]] < orderMap[nodes[j]]
		})
		order := 1
		explored := make(map[int]bool)
		for i := len(nodes) - 1; i >= 0; i-- {
			node := nodes[i]
			if !explored[node] {
				dfs(node, g, explored, &order, orderMap, leaderMap)
			}
		}
	}

	dfsLoop(gRev)
	dfsLoop(g)

	return leaderMap
}

func dfs(start int, g map[int][]int, explored map[int]bool, order *int, orderMap map[int]int, leaderMap map[int]int) {
	type item struct {
		node int
		end  bool
	}
	stack := []item{{start, false}}
	for len(stack) > 0 {
		it := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if it.end {
			orderMap[it.node] = *order
			*order++
			continue
		}
		explored[it.node] = true
		leaderMap[it.node] = start
		stack = append(stack, item{it.node, true})
		for _, other := range g[it.node] {
			if !explored[other] {
				stack = append(stack, item{other, false})
			}
		}
	}
}
