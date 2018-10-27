package meds

import (
	"container/heap"
)

type Meds struct {
	lo *maxHeap
	hi *minHeap
}

func NewMeds() Meds {
	return Meds{&maxHeap{intHeap{}}, &minHeap{intHeap{}}}
}

//Update adds n to the internal heaps and returns the current median.
func (m *Meds) Update(n int) int {
	if m.lo.Len() > 0 && n < m.lo.intHeap[0] {
		heap.Push(m.lo, n)
	} else {
		heap.Push(m.hi, n)
	}

	var from, to heap.Interface = m.lo, m.hi
	if m.lo.Len() < m.hi.Len() {
		from, to = m.hi, m.lo
	}

	total := m.lo.Len() + m.hi.Len()
	for (from.Len() - to.Len()) != total%2 {
		heap.Push(to, heap.Pop(from))
	}

	if total%2 == 0 {
		if m.lo.Len() > 0 {
			return m.lo.intHeap[0]
		} else {
			return m.hi.intHeap[0]
		}
	} else {
		if m.lo.Len() > m.hi.Len() {
			return m.lo.intHeap[0]
		} else {
			return m.hi.intHeap[0]
		}
	}
}
