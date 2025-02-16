package task0013

func romanToInt_02(s string) int {
	mapping := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var decimal int = mapping[s[len(s)-1]]
	for i := range len(s) - 1 {
		current, next := mapping[s[i]], mapping[s[i+1]]
		if current < next {
			decimal -= current
		} else {
			decimal += current
		}
	}
	return decimal
}
