package tsp

import "math"

const threshold = 1e-3

func approx(a, b float64) bool {
	return math.Abs(a-b) <= threshold
}
