package task0026

func removeDuplicates(nums []int) int {
	currI := 0
	for i := 1; i < len(nums); i++ {
		if nums[currI] != nums[i] {
			currI += 1
			nums[currI] = nums[i]
		}
	}

	return currI + 1
}
