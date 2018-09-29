package qsort

// Qsort sorts sort an int slice in-place with QuickSort algorithm and returns the comparison used.
// choosePivot should return the index of the pivot element it chosen.
func Qsort(a []int, choosePivot func([]int) int) int {
	n := len(a)

	if n <= 1 {
		return 0
	}

	pi := choosePivot(a)
	half1, half2 := partition(a, pi)

	count1 := Qsort(half1, choosePivot)
	count2 := Qsort(half2, choosePivot)

	return (n - 1) + count1 + count2
}

// First quick sorts an int slice with choosing always the first element as the pivot.
func First(a []int) int {
	return Qsort(a, func(a []int) int {
		return 0
	})
}

// Last quick sorts an int slice with choosing always the last element as the pivot.
func Last(a []int) int {
	return Qsort(a, func(a []int) int {
		return len(a) - 1
	})
}

// Median quick sorts an int slice with choosing always the median element of first, middle, last as the pivot.
func Median(a []int) int {
	return Qsort(a, func(a []int) int {
		n := len(a)
		fi, mi, li := 0, (n-1)/2, n-1
		f, m, l := a[fi], a[mi], a[li]

		if (m <= f && f <= l) || (l <= f && f <= m) {
			return fi
		}

		if (f <= m && m <= l) || (l <= m && m <= f) {
			return mi
		}

		if (f <= l && l <= m) || (m <= l && l <= f) {
			return li
		}

		panic("Unexpected")
	})
}

func partition(a []int, pi int) ([]int, []int) {
	n := len(a)
	a[0], a[pi] = a[pi], a[0]

	p := a[0]
	l, r := 1, 1
	for r < n {
		if a[r] < p {
			a[l], a[r] = a[r], a[l]
			l++
		}
		r++

	}

	a[0], a[l-1] = a[l-1], a[0]

	return a[:l-1], a[l:]
}
