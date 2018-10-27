package meds

import "testing"

func TestMeds(t *testing.T) {
	seqs := []struct {
		n, want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{6, 2},
		{5, 3},
		{4, 3},
		{7, 4},
	}
	m := NewMeds()

	numbers := []int{}
	for _, s := range seqs {
		numbers = append(numbers, s.n)
		if got := m.Update(s.n); got != s.want {
			t.Errorf("%v, median=%v, want %v", numbers, got, s.want)
		}
	}
}
