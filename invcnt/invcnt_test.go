package invcnt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		ar   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
		{[]int{2, 1}, 1},
		{[]int{3, 2, 1}, 3},
		{[]int{1, 3, 5, 2, 4, 6}, 3},
	}

	for _, test := range tests {
		if got := Count(test.ar); got != test.want {
			t.Errorf("Count(%v) = %d, want %d", test.ar, got, test.want)
		}
	}
}

func TestCountExample(t *testing.T) {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	reader := bufio.NewReader(f)

	var n int
	var ar []int
	for {
		_, err := fmt.Fscan(reader, &n)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		ar = append(ar, n)
	}

	got, want := Count(ar), 2407905288
	if got != want {
		t.Errorf("Example case failed")
	}
}
