package task0035

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1

		if nums[mid] >= target {
			right = mid - 1
		} else {
			if (mid == len(nums)-1) || nums[mid+1] >= target {
				return mid + 1
			}
			left = mid + 1
		}
	}

	return 0
}
