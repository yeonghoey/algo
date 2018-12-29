package tsp

import (
	"math"
)

// Vertex represents a position on a 2D plane.
type Vertex struct {
	X float64
	Y float64
}

type key struct {
	subset bitset
	end    int
}

// TSP calculates the cost of the minimum distance tour
// which visits all vertices once.
func TSP(vertices []Vertex) float64 {
	n := len(vertices)
	if n > 64 {
		panic("Not supported for vertices more than 64")
	}

	distTable := buildDistTable(vertices)
	A := make(map[key]float64)
	A[key{bitset(1), 0}] = 0
	for m := 1; m < n; m++ {
		subsets := listSubsets(n, m)
		for _, bs := range subsets {
			for j := 1; j < n; j++ {
				if !bs.contains(j) {
					continue
				}
				target := key{bs, j}
				bs1 := bs.unset(j)
				for k := 0; k < n; k++ {
					if !bs1.contains(k) {
						continue
					}
					distToK, ok := A[key{bs1, k}]
					if !ok {
						continue
					}

					dist1 := distToK + distTable[k][j]
					dist, ok := A[target]
					if !ok || dist1 < dist {
						A[target] = dist1
					}
				}
			}
		}
	}

	var all bitset
	for x := 0; x < n; x++ {
		all = all.set(x)
	}

	var dist *float64
	for j := 1; j < n; j++ {
		dist1 := A[key{all, j}] + distTable[j][0]
		if dist == nil || dist1 < *dist {
			dist = &dist1
		}
	}
	return *dist
}

func buildDistTable(vertices []Vertex) [][]float64 {
	calcDist := func(a, b Vertex) float64 {
		dx := a.X - b.X
		dy := a.Y - b.Y
		return math.Sqrt(dx*dx + dy*dy)
	}

	n := len(vertices)
	table := make([][]float64, n)
	for i, from := range vertices {
		row := make([]float64, n)
		for j, to := range vertices {
			row[j] = calcDist(from, to)
		}
		table[i] = row
	}
	return table
}

func listSubsets(n, numNonZero int) []bitset {
	bitsets := make([]bitset, 0)
	var f func(bitset, int)
	f = func(bs bitset, nz int) {
		if nz == 0 {
			bitsets = append(bitsets, bs)
			return
		}
		for x := 1; x < n; x++ {
			f(bs.set(x), nz-1)
		}
	}
	f(1, numNonZero)
	return bitsets
}
