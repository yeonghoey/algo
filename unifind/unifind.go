package unifind

// UniFind is a map from node number to its Union-Find meta data.
type UniFind map[int]*elem

type elem struct {
	node   int
	rank   int
	parent *elem
}

// Add adds a new node to UniFind as a root.
// ok will be false if the node already exists.
func (uf UniFind) Add(node int) (ok bool) {
	_, exists := uf[node]
	if !exists {
		uf[node] = &elem{node, 0, nil}
		ok = true
	}
	return
}

// Find finds the leader of the node.
// It also updates the parents of elements on the path.
func (uf UniFind) Find(node int) (leader int, ok bool) {
	this, exists := uf[node]
	if !exists {
		return 0, false
	}

	// Find the leader.
	trace := []*elem{}
	for this.parent != nil {
		trace = append(trace, this)
		this = this.parent
	}

	// Update the parents of the elements on the path to the leader.
	for _, e := range trace {
		e.parent = this
	}

	return this.node, true
}

// Merge merges nodeA's set with nodeB's set.
func (uf UniFind) Merge(nodeA, nodeB int) (ok bool) {
	leaderA, ok := uf.Find(nodeA)
	if !ok {
		return false
	}
	leaderB, ok := uf.Find(nodeB)
	if !ok {
		return false
	}

	// There is no need to union
	if leaderA == leaderB {
		return
	}

	a := uf[leaderA]
	b := uf[leaderB]
	if a.rank == b.rank {
		a.rank++
	}

	if a.rank < b.rank {
		a.parent = b
	} else {
		b.parent = a
	}

	return true
}
