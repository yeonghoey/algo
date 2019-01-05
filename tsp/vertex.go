package tsp

// Vertex represents a position on a 2D plane.
type Vertex struct {
	X float64
	Y float64
}

// Dist2 returns the squared distance from the vertex to another.
func (v *Vertex) Dist2(o Vertex) float64 {
	dx := v.X - o.X
	dy := v.Y - o.Y
	return dx*dx + dy*dy
}
