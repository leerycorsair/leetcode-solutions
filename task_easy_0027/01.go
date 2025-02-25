package task0027

func removeElement(nums []int, val int) int {
	deleted := 0
	for i := 0; i < len(nums)-deleted; {
		if nums[i] == val {
			nums[i] = nums[len(nums)-1-deleted]
			deleted += 1
		} else {
			i += 1
		}
	}

	return len(nums) - deleted
}
