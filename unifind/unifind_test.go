package unifind

import "testing"

func TestUniFind(t *testing.T) {
	uf := UniFind{}
	for i := 0; i < 5; i++ {
		uf.Add(i)
	}
	verify(t, uf, 0, 0)
	verify(t, uf, 1, 1)
	uf.Merge(0, 1)
	verify(t, uf, 0, 0)
	verify(t, uf, 1, 0)
	verify(t, uf, 2, 2)
	uf.Merge(0, 2)
	verify(t, uf, 2, 0)
	uf.Merge(3, 4)
	verify(t, uf, 4, 3)
	uf.Merge(2, 3)
	verify(t, uf, 3, 0)
	verify(t, uf, 4, 0)
}

func verify(t *testing.T, uf UniFind, node, want int) {
	leader, ok := uf.Find(node)
	if !(ok && leader == want) {
		t.Errorf("Find(%d) = (%d, %t), want (%d, %t)",
			node, leader, ok, want, true)
	}
}
