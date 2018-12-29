package tsp

import "math"

// Vertex represents a position on a 2D plane.
type Vertex struct {
	X float64
	Y float64
}

// TSP calculates the cost of the minimum distance tour
// which visits all vertices once.
func TSP(vertices []Vertex) float64 {
	n := len(vertices)
	if n > 64 {
		panic("Not supported for vertices more than 64")
	}

	distanceTable := buildDistanceTable(vertices)
	A := make(map[bitset]map[int]float64)
	for m := 1; m < n; m++ {
		subsets := listSubsets(n, m)
		for _, bs := range subsets {
			for j := 1; j < n; j++ {
				if !bs.contains(j) {
					continue
				}

				for k := 0; k < n; k++ {
					if k == j {
						continue
					}
				}
			}
		}
	}
	return 0.0
}

func buildDistanceTable(vertices []Vertex) [][]float64 {
	calcDistance := func(a, b Vertex) float64 {
		dx := a.X - b.X
		dy := a.Y - b.Y
		return math.Sqrt(dx*dx + dy*dy)
	}

	n := len(vertices)
	table := make([][]float64, n)
	for i, from := range vertices {
		row := make([]float64, n)
		for j, to := range vertices {
			row[j] = calcDistance(from, to)
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
