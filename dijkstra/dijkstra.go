package dijkstra

type Vertex = int

type Dist = int

type Edge struct {
	Tail   Vertex
	Head   Vertex
	Length Dist
}

func Dijkstra(edges []Edge, start Vertex) map[Vertex]Dist {
	processed := map[Vertex]bool{start: true}
	shortests := map[Vertex]Dist{start: 0}

	nPrevProcessed := 0
	for len(processed) > nPrevProcessed {
		nPrevProcessed = len(processed)

		best, bestE := -1, Edge{}
		for _, e := range edges {
			if !(processed[e.Tail] && !processed[e.Head]) {
				continue
			}
			criterion := shortests[e.Tail] + e.Length
			if best == -1 || criterion < best {
				best = criterion
				bestE = e
			}
		}

		if best != -1 {
			processed[bestE.Head] = true
			shortests[bestE.Head] = shortests[bestE.Tail] + bestE.Length
		}

	}
	return shortests
}
