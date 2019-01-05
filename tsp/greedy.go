package tsp

import "math"

// Greedy calculates a tour cost by picking the closest edge
// from the last vertex, staring from the first vertex.
func Greedy(vertices []Vertex) float64 {
	n := len(vertices)
	total := 0.0
	last := 0
	visited := make([]bool, n)
	for i := 1; i < n; i++ {
		visited[last] = true
		from := vertices[last]

		best2 := -1.0
		nextIdx := -1
		for idx, to := range vertices {
			if visited[idx] {
				continue
			}
			dist2 := from.Dist2(to)
			if best2 < 0 || dist2 < best2 {
				best2 = dist2
				nextIdx = idx
			}
		}
		total += math.Sqrt(best2)
		last = nextIdx
	}
	backTo0 := math.Sqrt(vertices[last].Dist2(vertices[0]))
	total += backTo0
	return total
}
