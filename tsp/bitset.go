package tsp

type bitset uint64

var flags [64]bitset

func init() {
	for i := range flags {
		flags[i] = 1 << bitset(i)
	}
}

func (bs bitset) set(x int) bitset {
	return bs | flags[x]
}

func (bs bitset) unset(x int) bitset {
	return bs & ^flags[x]
}

func (bs bitset) contains(x int) bool {
	return bs&flags[x] > 0
}
