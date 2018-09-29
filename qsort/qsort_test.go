package qsort

import (
	"reflect"
	"runtime"
	"sort"
	"testing"
)

func TestQsort(t *testing.T) {
	cases := []struct {
		a      []int
		first  int
		last   int
		median int
	}{
		{[]int{1}, 0, 0, 0},
		{[]int{1, 2}, 1, 1, 1},
		{[]int{2, 1}, 1, 1, 1},
		{[]int{1, 2, 3}, 3, 3, 2},
		{[]int{1, 3, 2}, 3, 2, 2},
		{[]int{2, 1, 3}, 2, 3, 2},
		{[]int{2, 3, 1}, 2, 3, 2},
		{[]int{3, 1, 2}, 3, 2, 2},
		{[]int{3, 2, 1}, 3, 3, 2},
	}

	for _, c := range cases {
		verify(t, First, c.a, c.first)
		verify(t, Last, c.a, c.last)
		verify(t, Median, c.a, c.median)
	}

}

func verify(t *testing.T, qsort func([]int) int, ca []int, want int) {
	a := make([]int, len(ca))
	copy(a, ca)

	name := funcName(qsort)
	count := qsort(a)

	if !sort.IntsAreSorted(a) {
		t.Errorf("%s(%v) did not sort properly. Got %v", name, ca, a)
	}
	if count != want {
		t.Errorf("First(%v) = %d, want %d", ca, count, want)
	}
}

func funcName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
