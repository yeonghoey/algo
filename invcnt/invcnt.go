package invcnt

// Count counts inversions using Divide and Conquer strategy.
func Count(ar []int) int {
	_, count := sortAndCount(ar)
	return count
}

func sortAndCount(ar []int) ([]int, int) {
	n := len(ar)
	if n <= 1 {
		return ar, 0
	}

	sorted1, count1 := sortAndCount(ar[:n/2])
	sorted2, count2 := sortAndCount(ar[n/2:])
	merged, splitCount := mergeAndCount(sorted1, sorted2)

	return merged, count1 + count2 + splitCount
}

func mergeAndCount(ar1, ar2 []int) ([]int, int) {
	n1, n2 := len(ar1), len(ar2)
	n := n1 + n2
	ar := make([]int, n)

	count := 0
	i, i1, i2 := 0, 0, 0

	for i < n && i1 < n1 && i2 < n2 {
		if ar1[i1] <= ar2[i2] {
			ar[i] = ar1[i1]
			i1++
		} else {
			count += n1 - i1
			ar[i] = ar2[i2]
			i2++
		}
		i++
	}

	for i1 < n1 {
		ar[i] = ar1[i1]
		i1++
		i++
	}

	for i2 < n2 {
		ar[i] = ar2[i2]
		i2++
		i++
	}

	return ar, count
}
