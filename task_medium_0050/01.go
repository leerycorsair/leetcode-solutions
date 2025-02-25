package task0050

func myPow_01(x float64, n int) float64 {
	abs := func(num int) int {
		if num > 0 {
			return num
		}
		return -num
	}

	pow2 := func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}

		curPow := 1
		absPow := abs(n)
		for curPow < absPow {
			x = x * x
			curPow *= 2
		}

		if n > 0 {
			return x
		}

		return 1 / x
	}

	sum2 := func(n int) []int {
		var result []int

		absN := abs(n)
		for absN > 0 {
			biggest2 := 1
			biggest2pow := 0
			for biggest2*2 <= absN {
				biggest2 *= 2
				biggest2pow += 1
			}
			absN -= biggest2

			if n >= 0 {
				result = append(result, biggest2)
			} else {
				result = append(result, -biggest2)
			}
		}

		return result
	}

	result := 1.0
	pows2 := sum2(n)
	for _, pow := range pows2 {
		result *= pow2(x, pow)
	}

	return result
}
