package task0020

func isValid(s string) bool {
	stack := []rune{}

	mapping := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	for _, r := range s {
		if r == '[' || r == '(' || r == '{' {
			stack = append(stack, r)
		} else {
			var value rune

			if len(stack) > 0 {
				value = stack[len(stack)-1]
			} else {
				return false
			}

			if mapping[value] != r {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
