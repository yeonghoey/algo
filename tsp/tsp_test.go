package tsp

import (
	"math"
	"testing"
)

const threshold = 1e-3

func TestTSP2(t *testing.T) {
	vertices := []Vertex{
		{0, 0},
		{1, 0},
	}
	want := 2.0
	if got := TSP(vertices); !approx(got, want) {
		t.Errorf("TSP(%v) = %.3f, want %.3f",
			vertices, got, want)
	}
}

func TestTSP4(t *testing.T) {
	vertices := []Vertex{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
	}
	want := 4.0
	if got := TSP(vertices); !approx(got, want) {
		t.Errorf("TSP(%v) = %.3f, want %.3f",
			vertices, got, want)
	}
}

func approx(a, b float64) bool {
	return math.Abs(a-b) <= threshold
}
