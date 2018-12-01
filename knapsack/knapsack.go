package knapsack

// Item represents the element to be put into the knapsack.
type Item struct {
	Value  int
	Weight int
}

// Knapsack calculates the optimal solution given the knapsack's size and items.
func Knapsack(size int, items []Item) int {
	d := make([][]int, len(items)+1)
	for i := range d {
		d[i] = make([]int, size+1)
	}

	for ii := 1; ii <= len(items); ii++ {
		this := items[ii-1]
		for wi := 1; wi <= size; wi++ {
			d[ii][wi] = d[ii-1][wi]
			if wi-this.Weight >= 0 {
				addThis := d[ii-1][wi-this.Weight] + this.Value
				d[ii][wi] = max(d[ii][wi], addThis)
			}
		}
	}

	return d[len(items)][size]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
