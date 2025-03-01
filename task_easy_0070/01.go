package task0070

func climbStairs(n int) int {
	next, secondNext := 0, 1
	for ; n > 0; n-- {
		next, secondNext = secondNext, next+secondNext
	}
	return secondNext
}
