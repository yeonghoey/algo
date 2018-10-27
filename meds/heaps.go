package meds

type intHeap []int

func (h intHeap) Len() int      { return len(h) }
func (h intHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type minHeap struct {
	intHeap
}

type maxHeap struct {
	intHeap
}

func (m minHeap) Less(i, j int) bool { return m.intHeap[i] < m.intHeap[j] }
func (m maxHeap) Less(i, j int) bool { return m.intHeap[i] > m.intHeap[j] }
