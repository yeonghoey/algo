package tsp

import (
	"math"
)

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

	distTable := buildDistTable(vertices)
	subsetTable := buildSubsetTable(n)

	A := make(map[bitset][]float64)

	row := newRow(n)
	row[0] = 0
	A[bitset(0).set(0)] = row

	for m := 1; m < n; m++ {
		for _, bs := range subsetTable[m] {
			var elems []int
			for e := 0; e < n; e++ {
				if bs.contains(e) {
					elems = append(elems, e)
				}
			}
			for _, j := range elems {
				if j == 0 {
					continue
				}

				bs1 := bs.unset(j)
				for _, k := range elems {
					if j == k {
						continue
					}
					rowToK, ok := A[bs1]
					if !ok {
						continue
					}
					if rowToK == nil {
						continue
					}
					distToK := rowToK[k]
					if distToK < 0 {
						continue
					}

					dist1 := distToK + distTable[k][j]
					rowDist, ok := A[bs]
					if !ok {
						rowDist = newRow(n)
						A[bs] = rowDist
					}
					if rowDist[j] < 0 || dist1 < rowDist[j] {
						rowDist[j] = dist1
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
		dist1 := A[all][j] + distTable[j][0]
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

func buildSubsetTable(n int) [][]bitset {
	table := make([][]bitset, n)
	var f func(bitset, int, int)
	f = func(bs bitset, from, nz int) {
		if nz >= n {
			return
		}
		for x := from; x < n; x++ {
			s := bs.set(x)
			table[nz] = append(table[nz], s)
			f(s, x+1, nz+1)
		}
	}
	f(bitset(0).set(0), 1, 1)
	return table
}

func newRow(n int) []float64 {
	row := make([]float64, n)
	for i := range row {
		row[i] = -1
	}
	return row
}
