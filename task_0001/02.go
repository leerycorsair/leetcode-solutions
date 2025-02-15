package task0001

func twoSum_02(nums []int, target int) []int {
	m := make(map[int]int)
	for i, value := range nums {
		m[value] = i
	}

	for j, value := range nums {
		if i, ok := m[target-value]; ok && i != j {
			return []int{i, j}
		}
	}

	return []int{}
}
