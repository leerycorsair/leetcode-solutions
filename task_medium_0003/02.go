package main

func lengthOfLongestSubstring_02(s string) int {
	maxLen := 0
	symbols := make(map[byte]int)
	left := 0

	for right := 0; right < len(s); right++ {
		if lastSeen, ok := symbols[s[right]]; ok && lastSeen >= left {
			left = lastSeen + 1
		}

		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}

		symbols[s[right]] = right
	}

	return maxLen
}
