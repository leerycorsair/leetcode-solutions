package task0013

func romanToInt_01(s string) int {
	mapping := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var decimal int
	for i := range s {
		if i != len(s)-1 {
			if mapping[s[i]] < mapping[s[i+1]] {
				decimal -= mapping[s[i]]
			} else {
				decimal += mapping[s[i]]
			}

		} else {
			decimal += mapping[s[i]]
		}
	}
	return decimal
}
