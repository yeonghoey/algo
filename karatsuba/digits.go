package karatsuba

import "fmt"

type digits = []int8

func toDigits(s string) (digits, error) {
	n := len(s)
	ds := make([]int8, n)

	for i, c := range s {
		d := int8(c - '0')
		if d < 0 || d > 9 {
			return nil, fmt.Errorf("%q contains non-digits char", s)
		}
		ds[n-i-1] = d
	}

	return ds, nil
}

func isEqual(a, b digits) bool {
	return true
}
