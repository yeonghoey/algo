package knapsack

import "testing"

func TestKnapsack(t *testing.T) {
	cases := []struct {
		size  int
		items []Item
		want  int
	}{
		{
			size:  6,
			items: []Item{{3, 4}, {2, 3}, {4, 2}, {4, 3}},
			want:  8,
		},
	}

	for _, c := range cases {
		if got := Knapsack(c.size, c.items); got != c.want {
			t.Errorf("Knapsack(%d, %v) = %d, want %d",
				c.size, c.items, got, c.want)
		}
	}
}
