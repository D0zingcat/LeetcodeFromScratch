package problem0001


func twoSumS1(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i; j < len(nums); j++ {
			if (nums[i] + nums[j] == target) {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumS2(nums []int, target int) []int {
	m := make(map[int](int), len(nums))
	for i, v := range nums {
		m[v] = i+1
	}
	for i, v := range nums {
		if (m[target-v] != 0 && i != m[target-v]-1) {
			return []int{i, m[target-v]-1}
		}
	}
	return nil
}


func twoSumS3(nums []int, target int) []int {
	m := make(map[int](int), len(nums))
	for i, v := range nums {
		if (m[target-v] != 0 && i != m[target-v]-1) {
			return []int{i, m[target-v]-1}
		}
		m[v] = i+1
	}
	return nil
}

func twoSumS4(nums []int, target int) []int {
	buff := make(map[int](int), len(nums))
	for i, v := range nums {
		val, ok := buff[target - v]
		if ok {
			return []int{i, val}
		}
		buff[v] = i
	}
	return nil
}