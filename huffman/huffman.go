package huffman

import (
	"sort"
)

const meta = -1

type node struct {
	label  int
	weight int
	left   *node
	right  *node
}

type byWeight []*node

func (a byWeight) Len() int           { return len(a) }
func (a byWeight) Less(i, j int) bool { return a[i].weight < a[j].weight }
func (a byWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Huffman returns an optimal codeword for indexes of the weights.
// The codewords are represented in string, containing only 0 and 1.
func Huffman(weights []int) []string {
	nodes := make([]*node, len(weights))
	metas := make([]*node, 0)

	for i, w := range weights {
		nodes[i] = &node{i, w, nil, nil}
	}
	sort.Sort(byWeight(nodes))

	for len(nodes)+len(metas) > 1 {
		ncand := nodes[:min(2, len(nodes))]
		mcand := metas[:min(2, len(metas))]
		a, b := pick2(ncand, mcand)
		for _, x := range []*node{a, b} {
			if len(nodes) > 0 && nodes[0] == x {
				nodes = nodes[1:]
			} else {
				metas = metas[1:]
			}
		}
		m := &node{meta, a.weight + b.weight, b, a}
		metas = append(metas, m)
	}

	return build(metas[0], len(weights))
}

func pick2(ncand, mcand []*node) (a, b *node) {
	cand := make([]*node, len(ncand)+len(mcand))
	i := 0
	for _, n := range ncand {
		cand[i] = n
		i++
	}
	for _, m := range mcand {
		cand[i] = m
		i++
	}
	sort.Sort(byWeight(cand))
	return cand[0], cand[1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func build(root *node, N int) []string {
	codemap := make([]string, N)
	var walk func(n *node, code string)
	walk = func(n *node, code string) {
		if n == nil {
			return
		}
		if n.label != meta {
			codemap[n.label] = code
		}

		walk(n.left, code+"0")
		walk(n.right, code+"1")
	}
	walk(root, "")
	return codemap
}
