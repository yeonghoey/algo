package karatsuba

import "fmt"

type digits []int8

func toDigits(s string) (digits, error) {
	n := len(s)
	ds := make(digits, n)

	for i, c := range s {
		d := int8(c - '0')
		if d < 0 || d > 9 {
			return nil, fmt.Errorf("%q contains non-digits char", s)
		}
		ds[n-i-1] = d
	}

	return ds, nil
}

func (ds digits) isEqual(other digits) bool {
	if (ds == nil) != (other == nil) {
		return false
	}

	if len(ds) != len(other) {
		return false
	}

	for i, d := range ds {
		if d != other[i] {
			return false
		}
	}
	return true
}

func (ds digits) add(other digits) digits {
	var n int

	if len(ds) < len(other) {
		n = len(other)
	} else {
		n = len(ds)
	}

	var carry int8
	// NOTE: As there may be a carry, set cap to n+1
	result := make(digits, n, n+1)
	for i := 0; i < n; i++ {
		sum := carry
		if i < len(ds) {
			sum += ds[i]
		}
		if i < len(other) {
			sum += other[i]
		}
		result[i] = sum % 10
		carry = sum / 10
	}

	if carry > 0 {
		result = append(result, carry)
	}
	return result
}

func (ds digits) number() string {
	n := len(ds)
	rs := make([]rune, n)
	for i, d := range ds {
		rs[n-1-i] = rune(d + '0')
	}
	return string(rs)
}
