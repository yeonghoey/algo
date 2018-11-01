package twosum

// TwoSum returns the number of distinct pairs in a which sums to t.
func TwoSum(a []int, t int) int {
	m := map[int]bool{}
	count := 0
	for _, n := range a {
		if m[t-n] {
			m[t-n] = false
			count++
		} else {
			m[n] = true
		}
	}
	return count
}
