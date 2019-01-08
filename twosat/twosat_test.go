package twosat

import "testing"

func TestTwoSat(t *testing.T) {
	cases := []struct {
		clauses []Clause
		want    bool
	}{
		{
			[]Clause{{1, 2}},
			true,
		},
		{
			[]Clause{{1, 2}, {1, -2}, {-1, -2}},
			true,
		},
		{
			[]Clause{{1, 2}, {1, -2}, {-1, 3}, {-1, -3}},
			false,
		},
		{
			[]Clause{{-1, 2}, {-2, -1}, {1, 1}},
			false,
		},
		{
			[]Clause{{-1, 2}, {2, 2}},
			true,
		},
		{
			[]Clause{
				{1, 2},
				{1, 2},
				{2, -3},
			}, true,
		},
		{
			[]Clause{
				{1, -2},
				{-1, 2},
				{-2, 4},
				{-2, -4},
				{2, 4},
				{2, -4},
			}, false,
		},
	}

	for _, c := range cases {
		if got := TwoSat(c.clauses); got != c.want {
			t.Errorf("TwoSet(%v) = %t, want %t",
				c.clauses, got, c.want)
		}
	}
}
