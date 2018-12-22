package flowar

// Edge represents an edge of a directed graph.
type Edge struct {
	Tail   int
	Head   int
	Length int
}

// PathMap represents the length of a path from a vertex to another.
type PathMap map[int]map[int]int

// Get gets the path length from a to b.
func (pm PathMap) Get(a, b int) (length int, ok bool) {
	fromA := pm[a]
	if fromA == nil {
		return
	}
	length, ok = fromA[b]
	return
}

// Set sets a path from a to b with length.
func (pm PathMap) Set(a, b int, length int) {
	fromA := pm[a]
	if fromA == nil {
		fromA = make(map[int]int)
		pm[a] = fromA
	}
	fromA[b] = length
}

// FloydWarshall returns All-pairs shortest paths calculated by Floyd-Warshall algorithm.
// It also returns negCycle, which represents whether or not a negative cylce exists in the graph.
func FloydWarshall(edges []Edge) (shortestPaths PathMap, negCycle bool) {
	// Init
	V := make(map[int]bool)
	A := make(PathMap)

	for _, e := range edges {
		V[e.Tail] = true
		V[e.Head] = true
		A.Set(e.Tail, e.Head, e.Length)
	}

	for v := range V {
		A.Set(v, v, 0)
	}

	// Run dynamic programming
	for k := range V {
		A1 := make(PathMap)
		for i := range V {
			for j := range V {
				cands := []int{}

				ij, ok0 := A.Get(i, j)
				if ok0 {
					cands = append(cands, ij)
				}
				ik, ok1 := A.Get(i, k)
				kj, ok2 := A.Get(k, j)
				if ok1 && ok2 {
					cands = append(cands, ik+kj)
				}

				for _, length := range cands {
					prev, ok := A1.Get(i, j)
					if !ok || (ok && length < prev) {
						A1.Set(i, j, length)
					}
				}
			}
		}
		A = A1
	}

	// Return
	shortestPaths = A
	negCycle = false
	for v := range V {
		length, ok := A.Get(v, v)
		if ok && length < 0 {
			negCycle = true
			break
		}
	}
	return
}
