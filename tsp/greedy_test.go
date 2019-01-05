package tsp

import (
	"testing"
)

func TestGreedy2(t *testing.T) {
	vertices := []Vertex{
		{0, 0},
		{1, 0},
	}
	want := 2.0
	if got := Greedy(vertices); !approx(got, want) {
		t.Errorf("Greedy(%v) = %.3f, want %.3f",
			vertices, got, want)
	}
}

func TestGreedy4(t *testing.T) {
	vertices := []Vertex{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
	}
	want := 4.0
	if got := Greedy(vertices); !approx(got, want) {
		t.Errorf("Greedy(%v) = %.3f, want %.3f",
			vertices, got, want)
	}
}
