package meds

import (
	"container/heap"
	"testing"
)

func TestHeaps(t *testing.T) {
	cases := []struct {
		heap heap.Interface
		want []int
	}{
		{&minHeap{intHeap{1, 4, 2}}, []int{1, 2, 3, 4}},
		{&maxHeap{intHeap{1, 4, 2}}, []int{4, 3, 2, 1}},
	}

	for _, c := range cases {
		heap.Init(c.heap)
		heap.Push(c.heap, 3)

		got := []int{}
		for c.heap.Len() > 0 {
			x := heap.Pop(c.heap)
			got = append(got, x.(int))
		}

		if !isEqual(got, c.want) {
			t.Errorf("Popped elements from minHeap = %v, want %v", got, c.want)
		}
	}
}

func isEqual(a, b []int) bool {
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
