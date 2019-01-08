package twosat

import (
	"github.com/yeonghoey/algo/scc"
)

// Clause represents Xa v Xb.
// if the value of A or B is negative,
// it represents (not Xa) or (not Xb) respectively.
type Clause struct {
	A, B int
}

// TwoSat determines whether or not an assignment exists that satisfies
// all the give clauses.
func TwoSat(clauses []Clause) bool {
	edges := make([]scc.Edge, 0, len(clauses)*2)
	for _, cl := range clauses {
		edges = append(edges, scc.Edge{-cl.A, cl.B})
		edges = append(edges, scc.Edge{-cl.B, cl.A})
	}

	leaderMap := scc.SCC2(edges)

	for node, leader := range leaderMap {
		leader1, ok := leaderMap[-node]
		if ok && (leader1 == leader) {
			return false
		}
	}
	return true
}
