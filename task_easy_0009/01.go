package task0009

import "fmt"

func isPalindrome_01(x int) bool {
	x_str := fmt.Sprint(x)

	for i := range len(x_str) / 2 {
		if x_str[i] != x_str[len(x_str)-i-1] {
			return false
		}
	}

	return true
}
