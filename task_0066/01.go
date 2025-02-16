package task0066

func plusOne(digits []int) []int {
	var prev = (digits[len(digits)-1] + 1) / 10
	digits[len(digits)-1] = (digits[len(digits)-1] + 1) % 10
	for i := 1; i < len(digits); i++ {
		if prev > 0 {
			prev = (digits[len(digits)-i-1] + 1) / 10
			digits[len(digits)-i-1] = (digits[len(digits)-i-1] + 1) % 10
		}
	}

	if prev > 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
