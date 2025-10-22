package main

func lengthOfLongestSubstring_01(s string) int {
	maxLen := 0

	for i := 0; i < len(s); i++ {
		symbols := make(map[byte]struct{})
		currentLen := 0

		for j := i; j < len(s); j++ {
			if _, ok := symbols[s[j]]; ok {
				break
			} else {
				symbols[s[j]] = struct{}{}
				currentLen += 1
			}

		}

		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}
