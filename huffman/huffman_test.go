package huffman

import "testing"

func TestHuffman(t *testing.T) {
	cases := []struct {
		weights []int
		want    []string
	}{
		{
			[]int{1, 2},
			[]string{"1", "0"},
		},
		{
			[]int{1, 2, 4},
			[]string{"11", "10", "0"},
		},
		{
			[]int{60, 25, 10, 5},
			[]string{"0", "10", "110", "111"},
		},
		{
			[]int{3, 2, 6, 8, 2, 6},
			[]string{"011", "0101", "11", "00", "0100", "10"},
		},
	}

	for _, c := range cases {
		if got := Huffman(c.weights); !isEqual(got, c.want) {
			t.Errorf("Huffman(%v) = %v, want %v", c.weights, got, c.want)
		}
	}
}

func isEqual(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
