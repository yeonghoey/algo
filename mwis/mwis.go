package mwis

// MWIS calculates a maximum-weight independent set.
func MWIS(weights []int) []int {
	N := len(weights)
	dp := make([]int, N)
	dp[0] = weights[0]
	dp[1] = weights[1]
	for i := 2; i < N; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+weights[i])
	}

	elems := make([]int, 0)
	i := N - 1
	for i > 1 {
		if dp[i-1] > dp[i-2]+weights[i] {
			i--
		} else {
			elems = append(elems, i)
			i -= 2
		}
	}
	// Add 0 or 1.
	elems = append(elems, i)

	reverse(elems)
	return elems
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
