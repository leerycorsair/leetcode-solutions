package task0001

func twoSum_03(nums []int, target int) []int {
	m := make(map[int]int)
	for i, value := range nums {
		if j, ok := m[target-value]; ok {
			return []int{i, j}
		}

		m[value] = i
	}

	return []int{}
}
