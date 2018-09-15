package karatsuba

// Karatsuba do multiplication using Karatsuba algorithm.
func Karatsuba(xs, ys string) string {
	x, err := toDigits(xs)
	if err != nil {
		panic(err)
	}
	y, err := toDigits(ys)
	if err != nil {
		panic(err)
	}
	return karatsuba(x, y).number()
}

func karatsuba(x, y digits) digits {
	var n int
	if len(x) > len(y) {
		n = len(x)
	} else {
		n = len(y)
	}

	x = x.ensureN(n)
	y = y.ensureN(n)

	if n == 0 {
		return digits{0}
	}

	if n == 1 {
		z := x[0] * y[0]
		d0 := z % 10
		d1 := z / 10
		return digits{d0, d1}.noZeros()
	}

	a, b := x[n/2:], x[:n/2]
	c, d := y[n/2:], y[:n/2]

	ac := karatsuba(a, c)
	bd := karatsuba(b, d)
	abcd := karatsuba(a.add(b), c.add(d))
	adbc := abcd.sub(ac).sub(bd)
	return ac.shift((n / 2) * 2).add(adbc.shift(n / 2)).add(bd).noZeros()
}
